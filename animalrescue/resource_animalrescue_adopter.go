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
			"first_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"phone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"gender": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"birthdate": {
				Type:     schema.TypeString,
				Required: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"country": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"city": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zipcode": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pet_preferences": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
