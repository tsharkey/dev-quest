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

func SelectOptFromMap[V any](m map[string]V, label string, searcher func(string, int) bool) (V, error) {
	names := Keys(m)

	// TODO: figure out how to return an empty generic value if possible
	choice, _, err := Select(label, names, searcher)
	if err != nil {
		log.Fatalf("Error selecting from map: %s", err)
	}

	return m[choice], nil
}

func Select(label string, items []string, searcher func(string, int) bool) (string, int, error) {
	if searcher == nil {
		searcher = func(input string, index int) bool {
			return strings.Contains(items[index], input)
		}
	}

	prompt := promptui.Select{
		Label:             label,
		Items:             items,
		Searcher:          searcher,
		StartInSearchMode: true,
	}

	ix, _, err := prompt.Run()
	if err != nil {
		return "", ix, err
	}

	return items[ix], ix, nil
}

func GetResponse(label string, defaultVal string, validate func(input string) error) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Default:  defaultVal,
		Validate: validate,
	}

	res, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return res, nil
}
