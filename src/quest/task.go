package quest

import (
	"bytes"
	"dev-quest/src/util"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/osteele/liquid"
	"github.com/pkg/browser"
	"github.com/spf13/viper"
)

type Tasks []*Task

type Task struct {
	// TODO: check if mapstructure supports omitempty, also setup validation with go playground validator
	Type        string `yaml:"type" mapstructure:"type"` // one of response, select, cmd, url, clone, confirm
	Action      string `yaml:"action" mapstructure:"action"`
	Name        string `yaml:"name" mapstructure:"name"`
	Description string `yaml:"description" mapstructure:"description"`
	Optional    bool   `yaml:"optional" mapstructure:"optional,"`
	Completed   bool   `yaml:"completed" mapstructure:"completed"`
	ConfigKey   string `yaml:"config_key" mapstructure:"config_key"`
	ConfigType  string `yaml:"config_type" mapstructure:"config_type"`
	Default     string `yaml:"default" mapstructure:"default"`
}

/******************************************************************************
 * Task Actions
 ******************************************************************************/

func (t *Task) Do() error {
	var err error

	log.Printf("Starting task: %s", t.Description)

	if t.Optional && !t.Completed {
		err = util.Confirm("Do you want to do this task")
		if err != nil {
			return nil
		}
	}

	switch t.Type {
	case "confirm":
		err = t.Confirm()
	case "cmd":
		err = t.Cmd()
	case "url":
		err = t.Url()
	case "config":
		err = t.Config()
	default:
		return fmt.Errorf("unknown task type: %s", t.Type)
	}

	if err == nil {
		t.Completed = true
	}

	return err
}

// TODO what should happen when the user says no
func (t *Task) Confirm() error {
	err := util.Confirm(t.Description)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) Cmd() error {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()

	resolved, err := resolveCommandTemplateVars(t.Action)
	if err != nil {
		return err
	}

	args := strings.Split(resolved, " ")
	cmd := exec.Command(args[0], args[1:]...)
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(stdErr)

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("%s\n%s", err, buf.String())
	}

	s.Stop()

	return nil
}

func (t *Task) Url() error {
	return browser.OpenURL(t.Action)
}

func (t *Task) Config() error {
	res, err := util.GetResponse(t.Description, t.Default, t.validateConfigResponse())
	if err != nil {
		return err
	}

	viper.Set("user_config."+t.ConfigKey, res)
	return viper.WriteConfig()
}

/******************************************************************************
 * Task(s) Utility
 ******************************************************************************/

func (ts *Tasks) IsComplete() bool {
	for _, t := range *ts {
		if !t.Completed {
			return false
		}
	}

	return true
}

func (ts *Tasks) GetAvailable() Tasks {
	available := Tasks{}

	for _, t := range *ts {
		if !t.Completed {
			available = append(available, t)
		}
	}

	return available
}

func resolveCommandTemplateVars(cmd string) (string, error) {
	if strings.Contains(cmd, "{{") {
		eng := liquid.NewEngine()
		template := cmd

		// get the user config from viper
		userConfig := viper.GetStringMap("user_config")

		out, err := eng.ParseAndRenderString(template, userConfig)
		if err != nil {
			return "", err
		}

		log.Printf("[DEBUG] resolved command template: %s", out)

		return out, nil
	}

	return cmd, nil
}

func (t *Task) validateConfigResponse() func(string) error {
	return func(input string) error {
		switch t.ConfigType {
		case "dir":
			if _, err := os.Stat(input); err != nil {
				return fmt.Errorf("value does not exist")
			}
		}

		return nil
	}
}
