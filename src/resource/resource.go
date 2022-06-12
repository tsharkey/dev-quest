package resource

import (
	"dev-quest/src/util"
	"fmt"

	"github.com/pkg/browser"
)

type Resources []Resource

type Resource struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	URL         string `yaml:"url"`
}

func (resources Resources) Display() error {
	_, ix, err := util.Select("Select a resource", resources.displayStrings(), nil)
	if err != nil {
		return err
	}

	browser.OpenURL(resources[ix].URL)

	return nil
}

func (resources Resources) displayStrings() []string {
	var displayStrings []string
	for _, resource := range resources {
		displayStrings = append(displayStrings, resource.DisplayString())
	}
	return displayStrings
}

func (resource *Resource) DisplayString() string {
	return fmt.Sprintf("%s - %s", resource.Name, resource.Description)
}
