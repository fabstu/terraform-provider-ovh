package ovh

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIpLoadbalancingTcpFarm_importBasic(t *testing.T) {
	displayName := acctest.RandomWithPrefix(test_prefix)
	config := fmt.Sprintf(testAccIpLoadbalancingTcpFarmConfig,
		os.Getenv("OVH_IPLB_SERVICE"), displayName, 12345, "all")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckIpLoadbalancing(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
			},
			{
				ResourceName:      "ovh_iploadbalancing_tcp_farm.testfarm",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccIpLoadbalancingTcpFarmImportId("ovh_iploadbalancing_tcp_farm.testfarm"),
			},
		},
	})
}

func testAccIpLoadbalancingTcpFarmImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		testfarm, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("ovh_ip_loadbalancing_tcp_farm not found: %s", resourceName)
		}
		return fmt.Sprintf(
			"%s/%s",
			testfarm.Primary.Attributes["service_name"],
			testfarm.Primary.Attributes["id"],
		), nil
	}
}
