package quest

type Quests map[string]*Quest

type Dependecies []string

type Quest struct {
	Description string      `yaml:"description" mapstructure:"description"`
	Tasks       Tasks       `yaml:"tasks" mapstructure:"tasks"`
	Completed   bool        `yaml:"completed" mapstructure:"completed"`
	DependsOn   Dependecies `yaml:"depends_on" mapstructure:"depends_on"`
}

func (q *Quest) IsComplete() bool {
	for _, t := range q.Tasks {
		task := t
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

func (q *Quests) IsComplete() bool {
	for _, q := range *q {
		if !q.IsComplete() {
			return false
		}
	}

	return true
}

func (q Quests) GetAvailable() Quests {
	return findAvailable(q)
}
