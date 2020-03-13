package main

import (
	"github.com/anGie44/terraform-provider-animalrescue/animalrescue" // change this to the import path of your provider
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: animalrescue.Provider})
}
