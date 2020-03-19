package animalrescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAnimalRescueAdoption() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAnimalRescueAdoptionRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"adopter_first_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"adopter_last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"adoptee_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"adoption_date": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceAnimalRescueAdoptionRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
