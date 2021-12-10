package main

import (
	"github.com/gnewbury1/terraform-provider-routeros/routeros"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: routeros.NewProvider})
}
