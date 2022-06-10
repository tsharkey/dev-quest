package quest

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

type QuestLog struct {
	logFile string `yaml:"-"`

	Resources Resources `yaml:"resources" mapstructure:"resources"`
	Quests    Quests    `yaml:"quests" mapstructure:"quests"`
}

func GetQuestLog() (*QuestLog, error) {
	ql := QuestLog{}
	err := viper.Unmarshal(&ql)
	if err != nil {
		return nil, err
	}

	return &ql, nil
}

func (ql *QuestLog) reload() error {
	return viper.Unmarshal(&ql)
}

func (ql *QuestLog) hasAvailableQuests() bool {
	if err := ql.reload(); err != nil {
		log.Fatalf("error reloading quest log: %v", err)
		return false
	}

	return len(ql.availableQuests()) > 0
}

func (ql *QuestLog) Grind() error {
	for ql.hasAvailableQuests() {
		availableQuests := ql.availableQuests()

		// build out the list of quest names the user has available
		questNames := make([]string, 0, len(availableQuests))
		for name := range availableQuests {
			questNames = append(questNames, name)
		}

		prompt := promptui.Select{
			Label: "Which quest would you like to do?",
			Items: questNames,
		}

		_, name, err := prompt.Run()
		if err != nil {
			return err
		}

		if quest, ok := availableQuests[name]; ok {
			err = quest.Do()
			if err != nil {
				return err
			}

			err = completeQuest(name)
			if err != nil {
				return err
			}
		}
	}

	return fmt.Errorf("no more quests available")
}

func (ql *QuestLog) availableQuests() map[string]Quest {
	availableQuests := map[string]Quest{}

questLoop:
	for name, quest := range ql.Quests {
		if quest.Completed {
			continue
		}

		if len(quest.DependsOn) > 0 {
			for _, dependency := range quest.DependsOn {
				if !ql.Quests[dependency].Completed {
					continue questLoop
				}
			}
		}

		availableQuests[name] = quest
	}

	return availableQuests
}
