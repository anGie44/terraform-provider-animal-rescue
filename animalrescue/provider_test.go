package animalrescue

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAdopter string = os.Getenv("ANIMALRESCUE_TEST_ADOPTER")
var testAdoptee string = os.Getenv("ANIMALRESCUE_TEST_ADOPTEE")

var testAccProviders map[string]terraform.ResourceProvider
var testAccProviderFactories func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"animalrescue": testAccProvider,
	}
	testAccProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"github": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {

	if v := os.Getenv("ANIMALRESCUE_TEST_ADOPTER"); v == "" {
		t.Fatal("ANIMALRESCUE_TEST_ADOPTER must be set for acceptance tests")
	}
	if v := os.Getenv("ANIMALRESCUE_TEST_ADOPTEE"); v == "" {
		t.Fatal("ANIMALRESCUE_TEST_ADOPTEE must be set for acceptance tests")
	}

}

func testRespondJson(responseBody string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte(responseBody)); err != nil {
			return
		}
	}
}
