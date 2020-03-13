package rescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRescueAdopter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRescueAdopterRead,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceRescueAdopterRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
