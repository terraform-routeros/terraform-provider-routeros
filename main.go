package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/vaerh/terraform-provider-routeros/routeros"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debug bool

	// https://github.com/hashicorp/terraform-docs-common/blob/main/website/docs/plugin/debugging.mdx
	// https://developer.hashicorp.com/terraform/plugin/debugging
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		ProviderAddr: "vaerh/routeros",
		ProviderFunc: routeros.NewProvider,
		Debug:        debug,
	})
}
