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
			"id": {
				Type: schema.TypeInt,
			},
			"adopter_first_name": {
				Type: schema.TypeString,
			},
			"adopter_last_name": {
				Type: schema.TypeString,
			},
			"adoptee_name": {
				Type: schema.TypeString,
			},
			"adoption_date": {
				Type: schema.TypeString,
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
