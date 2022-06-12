package util

import (
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
)

func Confirm(label string) error {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	res, err := prompt.Run()
	if err != nil {
		return err
	}

	if res != "y" {
		return fmt.Errorf("user did not confirm")
	}

	return nil
}

func SelectOpt[V any](m map[string]V, label string, searcher func(string, int) bool) (V, error) {
	names := Keys(m)

	if searcher == nil {
		searcher = func(input string, index int) bool {
			return strings.Contains(names[index], input)
		}
	}

	prompt := promptui.Select{
		Label:             label,
		Items:             names,
		Searcher:          searcher,
		StartInSearchMode: true,
	}

	ix, _, err := prompt.Run()
	if err != nil {
		// TODO how can we fit this with generics
		log.Fatalf("Something went wrong with your selection: %v", err)
		// return V, err
	}

	return m[names[ix]], nil
}
