package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNotificationGotifyResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationGotifyResourceConfig("resourceGotifyTest", 0) + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationGotifyResourceConfig("resourceGotifyTest", 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("lidarr_notification_gotify.test", "priority", "0"),
					resource.TestCheckResourceAttrSet("lidarr_notification_gotify.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationGotifyResourceConfig("resourceGotifyTest", 0) + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationGotifyResourceConfig("resourceGotifyTest", 5),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("lidarr_notification_gotify.test", "priority", "5"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "lidarr_notification_gotify.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationGotifyResourceConfig(name string, priority int) string {
	return fmt.Sprintf(`
	resource "lidarr_notification_gotify" "test" {
		on_grab               = false
		on_import_failure     = false
		on_upgrade            = false
		on_download_failure   = false
		on_release_import     = false
		on_health_issue       = false
		on_application_update = false
	  
		include_health_warnings = false
		name                    = "%s"
	  
		server = "http://gotify-server.net"
		app_token = "Token"
		priority = %d
	}`, name, priority)
}
