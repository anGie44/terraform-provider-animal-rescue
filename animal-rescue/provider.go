package rescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"rescue_adopter": dataSourceRescueAdopter(),
			"rescue_adoptee": dataSourceRescueAdoptee(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"rescue_adopter": resourceRescueAdopter(),
			"rescue_adoptee": resourceRescueAdoptee(),
		},
	}
}
