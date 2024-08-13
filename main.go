package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/mohrezfadaei/terraform-provider-marzban/marzban"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: marzban.Provider,
	})
}
