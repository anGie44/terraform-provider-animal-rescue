package animalrescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"animalrescue_adopter": dataSourceAnimalRescueAdopter(),
			"animalrescue_adoptee": dataSourceAnimalRescueAdoptee(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"animalrescue_adopter": resourceAnimalRescueAdopter(),
			"animalrescue_adoptee": resourceAnimalRescueAdoptee(),
		},
	}
}
