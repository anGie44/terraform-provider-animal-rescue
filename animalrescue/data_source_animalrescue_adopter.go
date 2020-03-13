package animalrescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAnimalRescueAdopter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAnimalRescueAdopterRead,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceAnimalRescueAdopterRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
