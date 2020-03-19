package animalrescue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"animalrescue_adopter":  dataSourceAnimalRescueAdopter(),
			"animalrescue_adoptee":  dataSourceAnimalRescueAdoptee(),
			"animalrescue_adoption": dataSourceAnimalRescueAdoption(),
			"animalrescue_petpref":  dataSourceAnimalRescuePetPref(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"animalrescue_adopter":  resourceAnimalRescueAdopter(),
			"animalrescue_adoptee":  resourceAnimalRescueAdoptee(),
			"animalrescue_adoption": resourceAnimalRescueAdoption(),
			"animalresuce_petpref":  resourceAnimalRescuePetPref(),
		},
	}
}
