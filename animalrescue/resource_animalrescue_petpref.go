package animalrescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAnimalRescuePetPref() *schema.Resource {
	return &schema.Resource{
		Create: resourceAnimalRescuePetPrefCreate,
		Read:   resourceAnimalRescuePetPrefRead,
		Update: resourceAnimalRescuePetPrefUpdate,
		Delete: resourceAnimalRescuePetPrefDelete,

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

func resourceAnimalRescuePetPrefCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescuePetPrefRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescuePetPrefUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAnimalRescuePetPrefDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
