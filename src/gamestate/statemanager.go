package gamestate

import (
	"dev-quest/src/quest"

	"github.com/spf13/viper"
)

type StateManager struct {
}

func NewStateManager() *StateManager {
	return &StateManager{}
}

func (sm *StateManager) SaveStoryLines(sls quest.StoryLines) error {
	viper.Set("storylines", sls)
	return viper.WriteConfig()
}

func (sm *StateManager) LoadStoryLines() (quest.StoryLines, error) {
	sls := new(quest.StoryLines)
	return *sls, viper.UnmarshalKey("storylines", sls)
}
