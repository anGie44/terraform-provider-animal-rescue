package animalrescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAnimalRescueAdoptee() *schema.Resource {
	return &schema.Resource{
		Create: resourceAnimalRescueAdopteeCreate,
		Read:   resourceAnimalRescueAdopteeRead,
		Update: resourceAnimalRescueAdopteeUpdate,
		Delete: resourceAnimalRescueAdopteeDelete,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAnimalRescueAdopteeCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescueAdopteeRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescueAdopteeUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescueAdopteeDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
