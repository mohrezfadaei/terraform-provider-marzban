package marzban

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccNode(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"marzban": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: `
resource "marzban_node" "example" {
  name = "example_node"
  address = "192.168.1.1"
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("marzban_node.example", "name", "example_node"),
					resource.TestCheckResourceAttr("marzban_node.example", "address", "192.168.1.1"),
				),
			},
		},
	})
}
