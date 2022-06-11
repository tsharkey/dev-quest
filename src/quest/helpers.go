package quest

import (
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
)

func confirm(label string) error {
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

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func selectOpt[V any](m map[string]V, label string, searcher func(string, int) bool) (V, error) {
	names := keys(m)

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

func findAvailable[V Actionable](m map[string]V) (map[string]V, error) {
	available := make(map[string]V)

mainloop:
	for k, v := range m {
		if v.IsComplete() {
			continue
		}

		if len(v.GetDependencies()) > 0 {
			for _, dep := range v.GetDependencies() {
				if !m[dep].IsComplete() {
					continue mainloop
				}
			}
		}

		available[k] = v
	}

	if len(available) == 0 {
		return nil, fmt.Errorf("no available quests")
	}

	return available, nil
}

func hasAvailable[K comparable, V Actionable](m map[K]V) bool {
	for _, v := range m {
		if !v.IsComplete() {
			return true
		}
	}
	return false
}
