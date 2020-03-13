package main

import (
	rescue "github.com/anGie44/terraform-provider-animal-rescue/animal-rescue" // change this to the import path of your provider
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: rescue.Provider})
}
