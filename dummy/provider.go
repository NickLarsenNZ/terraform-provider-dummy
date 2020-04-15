package dummy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{

		ResourcesMap: map[string]*schema.Resource{
			"dummy_file": resourceFile(),
		},

		Schema: map[string]*schema.Schema{
			"some_block": {
				Description: "Example Configuration Block",
				Optional:    true,
				Type:        schema.TypeMap,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"some_key": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
		ConfigureFunc: providerConfigure,
	}
}

type config struct {
	some_key string
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c := &config{}

	if some_block, ok := d.GetOk("some_block"); ok {
		raw_config := some_block.(map[string]interface{})
		c.some_key = raw_config["some_key"].(string)
	} else {
		fmt.Println("---> some_block wasn't defined")
	}

	return c, nil
}
