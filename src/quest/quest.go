package quest

import "log"

type Quests map[string]Quest

type Dependecies []string

type Quest struct {
	ID        string      `yaml:"id" mapstructure:"id"`
	QuestText string      `yaml:"quest_text" mapstructure:"quest_text"`
	Tasks     Tasks       `yaml:"tasks" mapstructure:"tasks"`
	Completed bool        `yaml:"completed" mapstructure:"completed"`
	DependsOn Dependecies `yaml:"depends_on" mapstructure:"depends_on"`
}

func (q *Quest) Do() error {
	for i, task := range q.Tasks {
		if !task.Completed {
			if task.Optional {
				err := confirm(task.Description)
				if err != nil {
					log.Printf("skipped optional task...")
					completeTask(q.ID, i)
					return nil
				}
			}

			err := task.Do()
			if err != nil {
				return err
			}

			completeTask(q.ID, i)
		}
	}

	return nil
}
