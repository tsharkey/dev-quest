package quest

type StoryLines map[string]*StoryLine

type StoryLine struct {
	DependsOn []string `yaml:"depends_on" mapstructure:"depends_on"`
	Completed bool     `yaml:"complete" mapstructure:"completed"`
	Quests    Quests   `yaml:"quests" mapstructure:"quests"`
}

func (sl *StoryLine) IsComplete() bool {
	for _, quest := range sl.Quests {
		if !quest.Completed {
			return false
		}
	}

	sl.Completed = true

	return true
}

func (sl *StoryLine) GetDependencies() []string {
	return sl.DependsOn
}

func (sls StoryLines) Select() (StoryLine, error) {
	return StoryLine{}, nil
}

func (sl *StoryLine) Do() error {
	if hasAvailable(sl.Quests) {
		available, err := findAvailable(sl.Quests)
		if err != nil {
			return err
		}

		choice, err := selectOpt(available, "Please choose a quest", nil)
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
