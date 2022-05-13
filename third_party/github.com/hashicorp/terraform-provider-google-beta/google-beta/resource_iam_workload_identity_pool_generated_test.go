// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIAMBetaWorkloadIdentityPool_iamWorkloadIdentityPoolBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckIAMBetaWorkloadIdentityPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPool_iamWorkloadIdentityPoolBasicExample(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id"},
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPool_iamWorkloadIdentityPoolBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "example" {
  provider                  = google-beta
  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
}
`, context)
}

func TestAccIAMBetaWorkloadIdentityPool_iamWorkloadIdentityPoolFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckIAMBetaWorkloadIdentityPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPool_iamWorkloadIdentityPoolFullExample(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id"},
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPool_iamWorkloadIdentityPoolFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "example" {
  provider                  = google-beta
  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
  display_name              = "Name of pool"
  description               = "Identity pool for automated test"
  disabled                  = true
}
`, context)
}

func testAccCheckIAMBetaWorkloadIdentityPoolDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_iam_workload_identity_pool" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{IAMBetaBasePath}}projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}")
			if err != nil {
				return err
			}

			res, err := sendRequest(config, "GET", "", url, config.userAgent, nil)
			if err != nil {
				return nil
			}

			if v := res["state"]; v == "DELETED" {
				return nil
			}

			return fmt.Errorf("IAMBetaWorkloadIdentityPool still exists at %s", url)
		}

		return nil
	}
}