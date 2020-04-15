package dummy

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileCreate,
		Read:   resourceFileRead,
		Update: resourceFileUpdate,
		Delete: resourceFileDelete,
		Exists: resourceFileExists,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path of the file we will store stuff in",
			},
			"firstline": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "What to put on the first line",
			},
			"secondline": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "What to put on the second line",
			},
			"mtime": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "File modified time (computed)",
			},
		},
	}
}

func resourceFileCreate(d *schema.ResourceData, meta interface{}) error {
	path := d.Get("path").(string)
	firstline := d.Get("firstline").(string)
	secondline := d.Get("secondline").(string)
	thirdline := fmt.Sprintf("some_key: %s\n", meta.(*config).some_key)
	contents := []byte(fmt.Sprintf("%s\n%s\n%s\n", firstline, secondline, thirdline))

	fmt.Printf("---> Contents:\n%s\n", contents)

	err := ioutil.WriteFile(path, contents, 0644)
	if err != nil {
		return err
	}

	checksum := sha1.Sum([]byte(contents))
	hash := hex.EncodeToString(checksum[:])
	d.SetId(hash)

	return nil
}

func resourceFileRead(d *schema.ResourceData, meta interface{}) error {
	// Because we have an Exists we can assume the file exists at this point
	path := d.Get("path").(string)

	// Read in the file contents
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Check the hash of the content against the resource ID. If it doesn't reset the ID (indicating a destroy before create)
	checksum := sha1.Sum([]byte(contents))
	if hex.EncodeToString(checksum[:]) != d.Id() {
		d.SetId("")
		return nil
	}

	return nil
}

// https://www.terraform.io/docs/plugins/provider.html#resource-data
func resourceFileUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Update not implemented")
}

func resourceFileDelete(d *schema.ResourceData, meta interface{}) error {
	path := d.Get("path").(string)
	return os.Remove(path)
}

func resourceFileExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	path := d.Get("path").(string)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}
