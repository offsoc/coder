package main

import (
	"testing"
	"time"

	"github.com/coder/coder/v2/coderd/coderdtest"
	"github.com/coder/coder/v2/codersdk"
	"github.com/coder/coder/v2/enterprise/coderd/coderdenttest"
	"github.com/coder/coder/v2/testutil"
)

func TestDebugProvisionerAvailability(t *testing.T) {
	ctx := testutil.Context(t, testutil.WaitLong)
	
	// Create a test environment
	client, _, api, owner := coderdenttest.NewWithAPI(t, &coderdenttest.Options{
		Options: &coderdtest.Options{
			IncludeProvisionerDaemon: true,
		},
	})
	
	t.Logf("Owner organization ID: %s", owner.OrganizationID)
	
	// List existing provisioner daemons
	daemons, err := client.ProvisionerDaemons(ctx)
	if err != nil {
		t.Fatalf("Failed to get provisioner daemons: %v", err)
	}
	
	t.Logf("Found %d existing provisioner daemons:", len(daemons))
	for i, daemon := range daemons {
		t.Logf("  Daemon %d: name=%s, org=%s, tags=%+v, lastSeen=%v, valid=%v", 
			i, daemon.Name, daemon.OrganizationID, daemon.Tags, daemon.LastSeenAt, daemon.LastSeenAt.Valid)
		if daemon.LastSeenAt.Valid {
			age := time.Since(daemon.LastSeenAt.Time)
			t.Logf("    Age: %v", age)
		}
	}
	
	// Create a fresh provisioner daemon
	provisionerTags := map[string]string{"owner": "", "scope": "organization"}
	t.Logf("Creating fresh provisioner daemon with tags: %+v", provisionerTags)
	
	_, err = api.AGPL.CreateInMemoryTaggedProvisionerDaemon(
		ctx,
		"debug-fresh-daemon",
		[]codersdk.ProvisionerType{codersdk.ProvisionerTypeEcho},
		provisionerTags,
	)
	if err != nil {
		t.Fatalf("Failed to create fresh provisioner daemon: %v", err)
	}
	
	// Wait a moment and check again
	time.Sleep(2 * time.Second)
	
	daemons, err = client.ProvisionerDaemons(ctx)
	if err != nil {
		t.Fatalf("Failed to get provisioner daemons after creation: %v", err)
	}
	
	t.Logf("Found %d provisioner daemons after creation:", len(daemons))
	for i, daemon := range daemons {
		t.Logf("  Daemon %d: name=%s, org=%s, tags=%+v, lastSeen=%v, valid=%v", 
			i, daemon.Name, daemon.OrganizationID, daemon.Tags, daemon.LastSeenAt, daemon.LastSeenAt.Valid)
		if daemon.LastSeenAt.Valid {
			age := time.Since(daemon.LastSeenAt.Time)
			t.Logf("    Age: %v", age)
		}
		
		if daemon.OrganizationID != owner.OrganizationID {
			t.Logf("    WARNING: Daemon organization %s does not match expected %s", 
				daemon.OrganizationID, owner.OrganizationID)
		}
	}
	
	// Create a simple template version to see its job tags
	version := coderdtest.CreateTemplateVersion(t, client, owner.OrganizationID, nil)
	coderdtest.AwaitTemplateVersionJobCompleted(t, client, version.ID)
	
	versionDetails, err := client.TemplateVersion(ctx, version.ID)
	if err != nil {
		t.Fatalf("Failed to get template version: %v", err)
	}
	
	t.Logf("Template version job details:")
	t.Logf("  Job ID: %s", versionDetails.Job.ID)
	t.Logf("  Job Tags: %+v", versionDetails.Job.Tags)
	t.Logf("  Job Status: %s", versionDetails.Job.Status)
	t.Logf("  Organization ID: %s", versionDetails.OrganizationID)
	
	t.Logf("\nSummary:")
	t.Logf("- Owner org: %s", owner.OrganizationID)
	t.Logf("- Template version org: %s", versionDetails.OrganizationID)
	t.Logf("- Template job tags: %+v", versionDetails.Job.Tags)
	t.Logf("- Fresh daemon tags: %+v", provisionerTags)
	
	// Check if any daemon matches the template job
	matchingDaemons := 0
	for _, daemon := range daemons {
		if daemon.OrganizationID == versionDetails.OrganizationID {
			// Check if tags match (simplified check)
			tagsMatch := true
			for key, value := range versionDetails.Job.Tags {
				if daemon.Tags[key] != value {
					tagsMatch = false
					break
				}
			}
			if tagsMatch {
				matchingDaemons++
				t.Logf("- Daemon '%s' matches template job (org + tags)", daemon.Name)
			} else {
				t.Logf("- Daemon '%s' in same org but tags don't match: daemon=%+v vs job=%+v", 
					daemon.Name, daemon.Tags, versionDetails.Job.Tags)
			}
		} else {
			t.Logf("- Daemon '%s' in different org: %s vs %s", 
				daemon.Name, daemon.OrganizationID, versionDetails.OrganizationID)
		}
	}
	
	t.Logf("\nResult: %d matching provisioner daemons found", matchingDaemons)
	if matchingDaemons == 0 {
		t.Logf("This explains why the test fails - no matching provisioners!")
	} else {
		t.Logf("Provisioners should be available - issue might be elsewhere")
	}
}
