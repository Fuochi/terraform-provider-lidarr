package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"lidarr": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("LIDARR_URL"); v == "" {
		t.Skip("LIDARR_URL must be set for acceptance tests")
	}
	if v := os.Getenv("LIDARR_API_KEY"); v == "" {
		t.Skip("LIDARR_API_KEY must be set for acceptance tests")
	}
}
