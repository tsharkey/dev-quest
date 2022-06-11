package quest

import (
	"github.com/spf13/viper"
)

type QuestLog struct {
	logFile string `yaml:"-"`

	Resources  Resources  `yaml:"resources" mapstructure:"resources"`
	StoryLines StoryLines `yaml:"storylines" mapstructure:"storylines"`
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

func (ql *QuestLog) Start() error {
	for hasAvailable(ql.StoryLines) {
		available, err := findAvailable(ql.StoryLines)
		if err != nil {
			return err
		}

		choice, err := selectOpt(available, "Please choose a story line", nil)
		if err != nil {
			return err
		}

		err = choice.Do()
		if err != nil {
			return err
		}
	}
	return nil
}
