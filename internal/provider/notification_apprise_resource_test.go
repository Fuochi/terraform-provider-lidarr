package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotificationAppriseResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationAppriseResourceConfig("resourceAppriseTest", "token123") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationAppriseResourceConfig("resourceAppriseTest", "token123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("lidarr_notification_apprise.test", "auth_password", "token123"),
					resource.TestCheckResourceAttrSet("lidarr_notification_apprise.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationAppriseResourceConfig("resourceAppriseTest", "token123") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationAppriseResourceConfig("resourceAppriseTest", "token234"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("lidarr_notification_apprise.test", "auth_password", "token234"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "lidarr_notification_apprise.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationAppriseResourceConfig(name, token string) string {
	return fmt.Sprintf(`
	resource "lidarr_notification_apprise" "test" {
		on_grab                 = false
		on_import_failure       = false
		on_upgrade              = false
		on_download_failure     = false
		on_release_import   	= false
		on_health_issue         = false
		on_application_update   = false
	  
		include_health_warnings = false
		name                    = "%s"
	  
		notification_type = 1
		server_url = "https://apprise.go"
		auth_username = "User"
		auth_password = "%s"
		field_tags = ["warning","skull"]
	}`, name, token)
}
