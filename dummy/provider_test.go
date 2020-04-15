package dummy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"testing"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProvider = Provider().(*schema.Provider)

	raw := map[string]interface{}{
		"some_block": map[string]interface{}{
			"some_key": "some_value",
		},
	}

	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(raw))
	if err != nil {
		fmt.Printf("provider config error: %s", err)
	}

	testAccProviders = map[string]terraform.ResourceProvider{
		"dummy": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

// // Ensure Environment variables are set for Acceptance Testing
// func testAccPreCheck(t *testing.T) {
// 	if v := os.Getenv("..."); v == "" {
// 		t.Fatal("... must be set for acceptance tests")
// 	}
// }
