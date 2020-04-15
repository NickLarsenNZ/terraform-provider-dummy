package dummy

import (
	"fmt"
	"os"
	"testing"

	"errors"
	"io/ioutil"

	r "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	path            = "/tmp/something.txt"
	firstline       = "alpha"
	secondline      = "bravo"
	thirdline       = "some_key: some_value" // todo: remove this after bug testing
	dummyFileConfig = `
		resource "dummy_file" "d" {
			path       = "%s"
			firstline  = "%s"
			secondline = "%s"
		}
	`
)

func TestDummyResourceFile(t *testing.T) {
	expectedContent := fmt.Sprintf("%s\n%s\n%s\n", firstline, secondline, thirdline)
	renderedConfig := fmt.Sprintf(dummyFileConfig, path, firstline, secondline)

	r.UnitTest(t, r.TestCase{
		PreCheck: func() {
			//testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []r.TestStep{
			{
				Config: renderedConfig,
				Check: func(s *terraform.State) error {
					content, err := ioutil.ReadFile(path)
					if err != nil {
						return err
						//return fmt.Errorf("was an error\ntemplate:\n%s\ngot:\n%s\nexpected:\n%s\n", renderedConfig, content, expectedContent)
					}
					if string(content) != expectedContent {
						return fmt.Errorf("template:\n%s\ngot:\n%s\nexpected:\n%s\n", renderedConfig, content, expectedContent)
					}
					return nil
				},
			},
		},
		CheckDestroy: func(*terraform.State) error {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				return nil
			}
			return errors.New("resource_dummy did not get destroyed")
		},
	})

}
