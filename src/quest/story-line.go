package quest

type Storylines map[string]*Storyline

type Storyline struct {
	DependsOn []string `yaml:"depends_on" mapstructure:"depends_on"`
	Completed bool     `yaml:"complete" mapstructure:"completed"`
	Quests    Quests   `yaml:"quests" mapstructure:"quests"`
}

func (sl *Storyline) IsComplete() bool {
	for _, q := range sl.Quests {
		quest := q
		if !quest.Completed {
			return false
		}
	}

	sl.Completed = true

	return true
}

func (sl *Storyline) GetDependencies() []string {
	return sl.DependsOn
}

func (sl Storylines) IsComplete() bool {
	for _, sl := range sl {
		if !sl.IsComplete() {
			return false
		}
	}

	return true
}

func (sl Storylines) GetAvailable() Storylines {
	return findAvailable(sl)
}
