package sdkv2provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"regexp"
	"testing"
)

func TestAccPingAccessAuthOidc(t *testing.T) {
	re := regexp.MustCompile(`^(6\.[0-9])`)
	if !re.MatchString(paVersion) {
		t.Skipf("This test only runs against PingAccess 6.0 and above, not: %s", paVersion)
	}
	resourceName := "pingaccess_auth_oidc.demo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckPingAccessAuthOidcDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAuthOidcConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateRuntimeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "foo"),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "true"),
					resource.TestCheckResourceAttr(resourceName, "sts_token_exchange_endpoint", "https://foo/bar"),
					resource.TestCheckResourceAttr(resourceName, "use_slo", "false"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "true"),
				),
			},
			{
				Config: testAccPingAccessPingFederateRuntimeConfig(os.Getenv("PINGFEDERATE_TEST_IP"), "bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingAccessPingFederateRuntimeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "bar"),
					resource.TestCheckResourceAttr(resourceName, "skip_hostname_verification", "true"),
					resource.TestCheckResourceAttr(resourceName, "sts_token_exchange_endpoint", "https://foo/bar"),
					resource.TestCheckResourceAttr(resourceName, "use_slo", "false"),
					resource.TestCheckResourceAttr(resourceName, "trusted_certificate_group_id", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_proxy", "true"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAuthOidcConfig() string {
	return fmt.Sprintf(`


	resource "pingaccess_authn_req_list" "acc_test" {
		name   = "acctest_authoidc"
		authn_reqs = [
			"acr1",
			"acr2",
    	]
    }


	resource "pingaccess_auth_oidc" "demo" {
		authn_req_list_id = pingaccess_authn_req_list.acc_test.id
    	enabled = true
  		use_slo = false
  		oidc_configuration {
           client_credentials {
			  client_id = "cheese"
              client_secret {
                  value = "top_secret" 
              }
              credentials_type = "SECRET"
           }
           scopes = [ "openid" ]
        }
	}`)
}

func testAccCheckPingAccessAuthOidcDestroy(state *terraform.State) error {
	return nil
}
