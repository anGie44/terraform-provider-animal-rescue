package animalrescue

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAnimalRescueAdopteeDataSource_basic(t *testing.T) {
	dsn := "data.animalrescue.adoptee.test"
	adopteeName := "Doggo"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAnimalRescueAdopteeDataSourceConfig(adopteeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "adoptee.#"),
					resource.TestCheckResourceAttr(dsn, "name", adopteeName),
				),
			},
		},
	})
}

func testAccCheckAnimalRescueAdopteeDataSourceConfig(name string) string {
	return fmt.Sprintf(`
data "animalrescue_adoptee" "test" {
  name      = "%s"
}
`, name)
}
