package quest

type StoryLines map[string]*StoryLine

type StoryLine struct {
	DependsOn []string `yaml:"depends_on" mapstructure:"depends_on"`
	Completed bool     `yaml:"complete" mapstructure:"completed"`
	Quests    Quests   `yaml:"quests" mapstructure:"quests"`
}

func (sl *StoryLine) IsComplete() bool {
	for _, q := range sl.Quests {
		quest := q
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

func (sl StoryLines) IsComplete() bool {
	for _, sl := range sl {
		if !sl.IsComplete() {
			return false
		}
	}

	return true
}

func (sl StoryLines) GetAvailable() StoryLines {
	return findAvailable(sl)
}
