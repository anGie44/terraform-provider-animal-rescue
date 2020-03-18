package animalrescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAnimalRescuePetPref() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAnimalRescuePetPrefRead,

		Schema: map[string]*schema.Schema{
			"breed": {
				Type:     schema.TypeString,
				Required: true,
			},
			"age": {
				Type:     schema.TypeString,
				Required: true,
			},
			"gender": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceAnimalRescuePetPrefRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
