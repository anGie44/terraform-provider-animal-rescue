package animalrescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAnimalRescueAdopter() *schema.Resource {
	return &schema.Resource{
		Create: resourceAnimalRescueAdopterCreate,
		Read:   resourceAnimalRescueAdopterRead,
		Update: resourceAnimalRescueAdopterUpdate,
		Delete: resourceAnimalRescueAdopterDelete,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAnimalRescueAdopterCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescueAdopterRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescueAdopterUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescueAdopterDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
