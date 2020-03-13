package animalrescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAnimalRescueAdoptee() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAnimalRescueAdopteeRead,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceAnimalRescueAdopteeRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
