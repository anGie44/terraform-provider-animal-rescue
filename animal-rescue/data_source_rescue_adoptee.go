package rescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRescueAdoptee() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRescueAdopteeRead,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceRescueAdopteeRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
