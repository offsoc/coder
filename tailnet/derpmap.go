package tailnet

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/xerrors"
	"tailscale.com/ipn/ipnstate"
	"tailscale.com/tailcfg"
)

const DisableSTUN = "disable"

func STUNRegions(baseRegionID int, stunAddrs []string) ([]*tailcfg.DERPRegion, error) {
	regions := make([]*tailcfg.DERPRegion, 0, len(stunAddrs))
	for index, stunAddr := range stunAddrs {
		if stunAddr == DisableSTUN {
			return []*tailcfg.DERPRegion{}, nil
		}

		host, rawPort, err := net.SplitHostPort(stunAddr)
		if err != nil {
			return nil, xerrors.Errorf("split host port for %q: %w", stunAddr, err)
		}
		port, err := strconv.Atoi(rawPort)
		if err != nil {
			return nil, xerrors.Errorf("parse port for %q: %w", stunAddr, err)
		}

		regionID := baseRegionID + index + 1
		regions = append(regions, &tailcfg.DERPRegion{
			EmbeddedRelay: false,
			RegionID:      regionID,
			RegionCode:    fmt.Sprintf("coder_stun_%d", regionID),
			RegionName:    fmt.Sprintf("Coder STUN %d", regionID),
			Nodes: []*tailcfg.DERPNode{{
				Name:     fmt.Sprintf("%dstun0", regionID),
				RegionID: regionID,
				HostName: host,
				STUNOnly: true,
				STUNPort: port,
			}},
		})
	}

	return regions, nil
}

// NewDERPMap constructs a DERPMap from a set of STUN addresses and optionally a remote
// URL to fetch a mapping from e.g. https://controlplane.tailscale.com/derpmap/default.
//
//nolint:revive
func NewDERPMap(ctx context.Context, region *tailcfg.DERPRegion, stunAddrs []string, remoteURL, localPath string, disableSTUN bool) (*tailcfg.DERPMap, error) {
	if remoteURL != "" && localPath != "" {
		return nil, xerrors.New("a remote URL or local path must be specified, not both")
	}
	if disableSTUN {
		stunAddrs = nil
	}

	// stunAddrs only applies when a default region is set. Each STUN node gets
	// it's own region ID because netcheck will only try a single STUN server in
	// each region before canceling the region's STUN check.
	addRegions := []*tailcfg.DERPRegion{}
	if region != nil {
		addRegions = append(addRegions, region)
		stunRegions, err := STUNRegions(region.RegionID, stunAddrs)
		if err != nil {
			return nil, xerrors.Errorf("create stun regions: %w", err)
		}
		addRegions = append(addRegions, stunRegions...)
	}

	derpMap := &tailcfg.DERPMap{
		Regions: map[int]*tailcfg.DERPRegion{},
	}
	if remoteURL != "" {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, remoteURL, nil)
		if err != nil {
			return nil, xerrors.Errorf("create request: %w", err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, xerrors.Errorf("get derpmap: %w", err)
		}
		defer res.Body.Close()
		err = json.NewDecoder(res.Body).Decode(&derpMap)
		if err != nil {
			return nil, xerrors.Errorf("fetch derpmap: %w", err)
		}
	}
	if localPath != "" {
		content, err := os.ReadFile(localPath)
		if err != nil {
			return nil, xerrors.Errorf("read derpmap from %q: %w", localPath, err)
		}
		err = json.Unmarshal(content, &derpMap)
		if err != nil {
			return nil, xerrors.Errorf("unmarshal derpmap: %w", err)
		}
	}

	// Add our custom regions to the DERP map.
	if len(addRegions) > 0 {
		for _, region := range addRegions {
			_, conflicts := derpMap.Regions[region.RegionID]
			if conflicts {
				return nil, xerrors.Errorf("a default region ID %d (%s - %q) conflicts with a remote region from %q", region.RegionID, region.RegionCode, region.RegionName, remoteURL)
			}
			derpMap.Regions[region.RegionID] = region
		}
	}

	// Remove all STUNPorts from DERPy nodes, and fully remove all STUNOnly
	// nodes.
	if disableSTUN {
		for _, region := range derpMap.Regions {
			newNodes := make([]*tailcfg.DERPNode, 0, len(region.Nodes))
			for _, node := range region.Nodes {
				node.STUNPort = -1
				if !node.STUNOnly {
					newNodes = append(newNodes, node)
				}
			}
			region.Nodes = newNodes
		}
	}

	// Fail if the DERP map has no regions or no DERP nodes.
	badDerpMapMsg := "A valid DERP map is required for networking to work. You must either supply your own DERP map or use the built-in DERP server"
	if len(derpMap.Regions) == 0 {
		return nil, xerrors.New("DERP map has no regions. " + badDerpMapMsg)
	}
	foundValidNode := false
regionLoop:
	for _, region := range derpMap.Regions {
		for _, node := range region.Nodes {
			if !node.STUNOnly {
				foundValidNode = true
				break regionLoop
			}
		}
	}
	if !foundValidNode {
		return nil, xerrors.New("DERP map has no DERP nodes. " + badDerpMapMsg)
	}

	return derpMap, nil
}

func ExtractPreferredDERPName(pingResult *ipnstate.PingResult, node *Node, derpMap *tailcfg.DERPMap) string {
	// Sometimes the preferred DERP doesn't match the one we're actually
	// connected with. Perhaps because the agent prefers a different DERP and
	// we're using that server instead.
	preferredDerpID := node.PreferredDERP
	if pingResult.DERPRegionID != 0 {
		preferredDerpID = pingResult.DERPRegionID
	}
	preferredDerp, ok := derpMap.Regions[preferredDerpID]
	preferredDerpName := fmt.Sprintf("Unnamed %d", preferredDerpID)
	if ok {
		preferredDerpName = preferredDerp.RegionName
	}

	return preferredDerpName
}

// ExtractDERPLatency extracts a map of derp region names to their latencies
func ExtractDERPLatency(node *Node, derpMap *tailcfg.DERPMap) map[string]time.Duration {
	latencyMs := make(map[string]time.Duration)

	// Convert DERP region IDs to friendly names for display in the UI.
	for rawRegion, latency := range node.DERPLatency {
		regionParts := strings.SplitN(rawRegion, "-", 2)
		regionID, err := strconv.Atoi(regionParts[0])
		if err != nil {
			continue
		}
		region, found := derpMap.Regions[regionID]
		if !found {
			// It's possible that a workspace agent is using an old DERPMap
			// and reports regions that do not exist. If that's the case,
			// report the region as unknown!
			region = &tailcfg.DERPRegion{
				RegionID:   regionID,
				RegionName: fmt.Sprintf("Unnamed %d", regionID),
			}
		}
		latencyMs[region.RegionName] = time.Duration(latency * float64(time.Second))
	}
	return latencyMs
}

// CompareDERPMaps returns true if the given DERPMaps are equivalent. Ordering
// of slices is ignored.
//
// If the first map is nil, the second map must also be nil for them to be
// considered equivalent. If the second map is nil, the first map can be any
// value and the function will return true.
func CompareDERPMaps(a *tailcfg.DERPMap, b *tailcfg.DERPMap) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return true
	}
	if len(a.Regions) != len(b.Regions) {
		return false
	}
	if a.OmitDefaultRegions != b.OmitDefaultRegions {
		return false
	}

	for id, region := range a.Regions {
		other, ok := b.Regions[id]
		if !ok {
			return false
		}
		if !compareDERPRegions(region, other) {
			return false
		}
	}
	return true
}

func compareDERPRegions(a *tailcfg.DERPRegion, b *tailcfg.DERPRegion) bool {
	if a == nil || b == nil {
		return false
	}
	if a.EmbeddedRelay != b.EmbeddedRelay {
		return false
	}
	if a.RegionID != b.RegionID {
		return false
	}
	if a.RegionCode != b.RegionCode {
		return false
	}
	if a.RegionName != b.RegionName {
		return false
	}
	if a.Avoid != b.Avoid {
		return false
	}
	if len(a.Nodes) != len(b.Nodes) {
		return false
	}

	// Convert both slices to maps so ordering can be ignored easier.
	aNodes := map[string]*tailcfg.DERPNode{}
	for _, node := range a.Nodes {
		aNodes[node.Name] = node
	}
	bNodes := map[string]*tailcfg.DERPNode{}
	for _, node := range b.Nodes {
		bNodes[node.Name] = node
	}

	for name, aNode := range aNodes {
		bNode, ok := bNodes[name]
		if !ok {
			return false
		}

		if aNode.Name != bNode.Name {
			return false
		}
		if aNode.RegionID != bNode.RegionID {
			return false
		}
		if aNode.HostName != bNode.HostName {
			return false
		}
		if aNode.CertName != bNode.CertName {
			return false
		}
		if aNode.IPv4 != bNode.IPv4 {
			return false
		}
		if aNode.IPv6 != bNode.IPv6 {
			return false
		}
		if aNode.STUNPort != bNode.STUNPort {
			return false
		}
		if aNode.STUNOnly != bNode.STUNOnly {
			return false
		}
		if aNode.DERPPort != bNode.DERPPort {
			return false
		}
		if aNode.InsecureForTests != bNode.InsecureForTests {
			return false
		}
		if aNode.ForceHTTP != bNode.ForceHTTP {
			return false
		}
		if aNode.STUNTestIP != bNode.STUNTestIP {
			return false
		}
	}

	return true
}
