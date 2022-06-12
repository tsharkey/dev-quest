package quest

import (
	"bytes"
	"dev-quest/src/userconfig"
	"dev-quest/src/util"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"github.com/pkg/browser"
	"github.com/spf13/viper"
	"gopkg.in/osteele/liquid.v1"
)

type Tasks []*Task

type Task struct {
	Type        string `yaml:"type" mapstructure:"type"` // one of response, select, cmd, url, clone, confirm
	Action      string `yaml:"action" mapstructure:"action"`
	Description string `yaml:"description" mapstructure:"description"`
	Optional    bool   `yaml:"optional" mapstructure:"optional"`
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

	if t.Optional && !t.Completed {
		err = util.Confirm("Do you want to do this task?")
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

	resolved := resolveCommandTemplateVars(t.Action)
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
	prompt := promptui.Prompt{
		Label:   "Enter value for " + t.ConfigKey,
		Default: t.Default,
		Validate: func(input string) error {
			switch t.ConfigType {
			case "dir":
				if _, err := os.Stat(input); err != nil {
					return fmt.Errorf("value does not exist")
				}
			}

			return nil
		},
	}

	value, err := prompt.Run()
	if err != nil {
		return err
	}

	return userconfig.Set(t.ConfigKey, value)
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

// TODO: clean this up- no panics
func resolveCommandTemplateVars(cmd string) string {
	if strings.Contains(cmd, "{{") {
		availbleBindings := map[string]interface{}{}
		eng := liquid.NewEngine()
		template := cmd
		err := viper.UnmarshalKey("user_config", &availbleBindings)
		if err != nil {
			panic(err)
		}

		out, err := eng.ParseAndRenderString(template, availbleBindings)
		if err != nil {
			panic(err)
		}

		log.Printf("[DEBUG] resolved command template: %s", out)

		return out
	}

	return cmd
}
