package quest

type Quests map[string]*Quest

type Dependecies []string

type Quest struct {
	ID        string      `yaml:"id" mapstructure:"id"`
	QuestText string      `yaml:"quest_text" mapstructure:"quest_text"`
	Tasks     Tasks       `yaml:"tasks" mapstructure:"tasks"`
	Completed bool        `yaml:"completed" mapstructure:"completed"`
	DependsOn Dependecies `yaml:"depends_on" mapstructure:"depends_on"`
}

func (q *Quest) IsComplete() bool {
	for _, task := range q.Tasks {
		if !task.Completed {
			return false
		}
	}

	q.Completed = true

	return true
}

func (q *Quest) GetDependencies() []string {
	return q.DependsOn
}

func (q *Quest) Do() error {
	if !q.Tasks.Done() {
		for _, task := range q.Tasks {
			err := task.Do()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
