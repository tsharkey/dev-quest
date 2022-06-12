package game

import (
	"dev-quest/src/gamestate"
	"dev-quest/src/quest"
	"dev-quest/src/util"
	"fmt"
)

type Game struct {
	StateManager *gamestate.StateManager
	StoryLines   quest.StoryLines
}

func NewGame(sm *gamestate.StateManager) *Game {
	return &Game{
		StateManager: sm,
	}
}

func (g *Game) init() error {
	var err error
	g.StoryLines, err = g.StateManager.LoadStoryLines()
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
	for availableStoryLines := g.StoryLines.GetAvailable(); len(availableStoryLines) > 0; availableStoryLines = g.StoryLines.GetAvailable() {
		chosenStoryLine, err := util.SelectOpt(availableStoryLines, "Please choose a story line to start:", nil)
		if err != nil {
			return err
		}

		// gets available quests and refreshes at every iteration
		for availableQuests := chosenStoryLine.Quests.GetAvailable(); len(availableQuests) > 0; availableQuests = chosenStoryLine.Quests.GetAvailable() {
			chosenQuest, err := util.SelectOpt(availableQuests, "Please choose a quest to start:", nil)
			if err != nil {
				return err
			}

			for !chosenQuest.IsComplete() {
				availableTasks := chosenQuest.Tasks.GetAvailable()
				for _, task := range availableTasks {
					if err := task.Do(); err != nil {
						return err
					}

					task.Completed = true

					if err := g.Save(); err != nil {
						return err
					}

				}

				chosenQuest.Completed = true

				if err := g.Save(); err != nil {
					return err
				}
			}
		}

		chosenStoryLine.Completed = true

		if err := g.Save(); err != nil {
			return err
		}
	}

	return fmt.Errorf("game complete")
}

func (g *Game) Save() error {
	err := g.StateManager.SaveStoryLines(g.StoryLines)
	if err != nil {
		return err
	}

	return nil
}
