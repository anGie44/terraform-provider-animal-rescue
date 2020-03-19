package animalrescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAnimalRescueAdoption() *schema.Resource {
	return &schema.Resource{
		Create: resourceAnimalRescueAdoptionCreate,
		Read:   resourceAnimalRescueAdoptionRead,
		Update: resourceAnimalRescueAdoptionUpdate,
		Delete: resourceAnimalRescueAdoptionDelete,

		Schema: map[string]*schema.Schema{
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

func resourceAnimalRescueAdoptionCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescueAdoptionRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescueAdoptionUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescueAdoptionDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
