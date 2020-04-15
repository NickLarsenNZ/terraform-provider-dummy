package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/nicklarsennz/terraform-provider-dummy/dummy"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dummy.Provider,
	})
}
