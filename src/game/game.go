package game

import (
	"dev-quest/src/quest"
	"dev-quest/src/resource"
	"dev-quest/src/util"
	"fmt"
)

type Game struct {
	Name        string             `yaml:"name" mapstructure:"name"`
	Description string             `yaml:"description" mapstructure:"description"`
	Storylines  quest.Storylines   `yaml:"storylines" mapstructure:"storylines"`
	Resources   resource.Resources `yaml:"resources" mapstructure:"resources"`
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) init() error {
	var err error
	g.Storylines, err = LoadStorylines()
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Start() error {
	err := g.init()
	if err != nil {
		return err
	}

	fmt.Println("Welcome to Dev Quest!")
	fmt.Println("")

	// gets available storylines and refreshes at every iteration
	for availableStorylines := g.Storylines.GetAvailable(); len(availableStorylines) > 0; availableStorylines = g.Storylines.GetAvailable() {
		chosenStoryline, err := util.SelectOptFromMap(availableStorylines, "Please choose a story line to start", nil)
		if err != nil {
			return err
		}

		// gets available quests and refreshes at every iteration
		for availableQuests := chosenStoryline.Quests.GetAvailable(); len(availableQuests) > 0; availableQuests = chosenStoryline.Quests.GetAvailable() {

			chosenQuest, err := util.SelectOptFromMap(availableQuests, "Please choose a quest to start", nil)
			if err != nil {
				return err
			}

			choice, _, err := util.Select("What would you like to do", []string{"start", "skip", "quit"}, nil)
			if err != nil {
				return err
			}

			switch choice {
			case "start":
				// nothing to do
			case "skip":
				chosenQuest.Completed = true
			case "quit":
				return nil
			}

			for !chosenQuest.IsComplete() {
				availableTasks := chosenQuest.Tasks.GetAvailable()
				for _, task := range availableTasks {
					if err := task.Do(); err != nil {
						return err
					}

					task.Completed = true

					if err := SaveStorylines(g.Storylines); err != nil {
						return err
					}

				}

				chosenQuest.Completed = true

				if err := SaveStorylines(g.Storylines); err != nil {
					return err
				}
			}
		}

		chosenStoryline.Completed = true

		if err := SaveStorylines(g.Storylines); err != nil {
			return err
		}
	}

	// TODO add some error types
	return fmt.Errorf("game complete")
}
