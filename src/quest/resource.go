package quest

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/browser"
)

type Resources []Resource

type Resource struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	URL         string `yaml:"url"`
}

func (resources Resources) Display() error {
	prompt := promptui.Select{
		Label: "Select a resource",
		Items: resources.displayStrings(),
		Searcher: func(input string, index int) bool {
			return strings.Contains(resources[index].Name, input)
		},
		StartInSearchMode: true,
	}

	ix, _, err := prompt.Run()
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
