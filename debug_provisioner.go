package main

import (
	"context"
	"log"
	"time"

	"github.com/coder/coder/v2/coderd/coderdtest"
	"github.com/coder/coder/v2/codersdk"
	"github.com/coder/coder/v2/enterprise/coderd/coderdenttest"
)

func main() {
	ctx := context.Background()
	
	// Create a test environment
	client, _, api, owner := coderdenttest.NewWithAPI(nil, &coderdenttest.Options{
		Options: &coderdtest.Options{
			IncludeProvisionerDaemon: true,
		},
	})
	
	log.Printf("Owner organization ID: %s", owner.OrganizationID)
	
	// List existing provisioner daemons
	daemons, err := client.ProvisionerDaemons(ctx)
	if err != nil {
		log.Fatalf("Failed to get provisioner daemons: %v", err)
	}
	
	log.Printf("Found %d existing provisioner daemons:", len(daemons))
	for i, daemon := range daemons {
		log.Printf("  Daemon %d: name=%s, org=%s, tags=%+v, lastSeen=%v, valid=%v", 
			i, daemon.Name, daemon.OrganizationID, daemon.Tags, daemon.LastSeenAt, daemon.LastSeenAt.Valid)
		if daemon.LastSeenAt.Valid {
			age := time.Since(daemon.LastSeenAt.Time)
			log.Printf("    Age: %v", age)
		}
	}
	
	// Create a fresh provisioner daemon
	provisionerTags := map[string]string{"owner": "", "scope": "organization"}
	log.Printf("Creating fresh provisioner daemon with tags: %+v", provisionerTags)
	
	_, err = api.AGPL.CreateInMemoryTaggedProvisionerDaemon(
		ctx,
		"debug-fresh-daemon",
		[]codersdk.ProvisionerType{codersdk.ProvisionerTypeEcho},
		provisionerTags,
	)
	if err != nil {
		log.Fatalf("Failed to create fresh provisioner daemon: %v", err)
	}
	
	// Wait a moment and check again
	time.Sleep(2 * time.Second)
	
	daemons, err = client.ProvisionerDaemons(ctx)
	if err != nil {
		log.Fatalf("Failed to get provisioner daemons after creation: %v", err)
	}
	
	log.Printf("Found %d provisioner daemons after creation:", len(daemons))
	for i, daemon := range daemons {
		log.Printf("  Daemon %d: name=%s, org=%s, tags=%+v, lastSeen=%v, valid=%v", 
			i, daemon.Name, daemon.OrganizationID, daemon.Tags, daemon.LastSeenAt, daemon.LastSeenAt.Valid)
		if daemon.LastSeenAt.Valid {
			age := time.Since(daemon.LastSeenAt.Time)
			log.Printf("    Age: %v", age)
		}
		
		if daemon.OrganizationID != owner.OrganizationID {
			log.Printf("    WARNING: Daemon organization %s does not match expected %s", 
				daemon.OrganizationID, owner.OrganizationID)
		}
	}
	
	// Create a simple template version to see its job tags
	version := coderdtest.CreateTemplateVersion(nil, client, owner.OrganizationID, nil)
	coderdtest.AwaitTemplateVersionJobCompleted(nil, client, version.ID)
	
	versionDetails, err := client.TemplateVersion(ctx, version.ID)
	if err != nil {
		log.Fatalf("Failed to get template version: %v", err)
	}
	
	log.Printf("Template version job details:")
	log.Printf("  Job ID: %s", versionDetails.Job.ID)
	log.Printf("  Job Tags: %+v", versionDetails.Job.Tags)
	log.Printf("  Job Status: %s", versionDetails.Job.Status)
	log.Printf("  Organization ID: %s", versionDetails.OrganizationID)
	
	log.Printf("\nSummary:")
	log.Printf("- Owner org: %s", owner.OrganizationID)
	log.Printf("- Template version org: %s", versionDetails.OrganizationID)
	log.Printf("- Template job tags: %+v", versionDetails.Job.Tags)
	log.Printf("- Fresh daemon tags: %+v", provisionerTags)
	
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
				log.Printf("- Daemon '%s' matches template job (org + tags)", daemon.Name)
			} else {
				log.Printf("- Daemon '%s' in same org but tags don't match: daemon=%+v vs job=%+v", 
					daemon.Name, daemon.Tags, versionDetails.Job.Tags)
			}
		} else {
			log.Printf("- Daemon '%s' in different org: %s vs %s", 
				daemon.Name, daemon.OrganizationID, versionDetails.OrganizationID)
		}
	}
	
	log.Printf("\nResult: %d matching provisioner daemons found", matchingDaemons)
	if matchingDaemons == 0 {
		log.Printf("This explains why the test fails - no matching provisioners!")
	} else {
		log.Printf("Provisioners should be available - issue might be elsewhere")
	}
}
