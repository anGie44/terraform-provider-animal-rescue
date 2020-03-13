package rescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRescueAdoptee() *schema.Resource {
	return &schema.Resource{
		Create: resourceRescueAdopteeCreate,
		Read:   resourceRescueAdopteeRead,
		Update: resourceRescueAdopteeUpdate,
		Delete: resourceRescueAdopteeDelete,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRescueAdopteeCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRescueAdopteeRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRescueAdopteeUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRescueAdopteeDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
