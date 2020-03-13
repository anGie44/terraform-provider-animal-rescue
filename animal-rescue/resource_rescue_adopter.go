package rescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRescueAdopter() *schema.Resource {
	return &schema.Resource{
		Create: resourceRescueAdopterCreate,
		Read:   resourceRescueAdopterRead,
		Update: resourceRescueAdopterUpdate,
		Delete: resourceRescueAdopterDelete,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRescueAdopterCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRescueAdopterRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRescueAdopterUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRescueAdopterDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
