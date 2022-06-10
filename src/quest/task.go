package quest

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/pkg/browser"
)

type Tasks []Task

type Task struct {
	Type        string `yaml:"type" mapstructure:"type"` // one of response, select, cmd, url, clone, confirm
	Action      string `yaml:"action" mapstructure:"action"`
	Description string `yaml:"description" mapstructure:"description"`
	Optional    bool   `yaml:"optional" mapstructure:"optional"`
	Completed   bool   `yaml:"completed" mapstructure:"completed"`
}

func (t *Task) Do() error {
	switch t.Type {
	case "confirm":
		return t.Confirm()
	case "cmd":
		return t.Cmd()
	case "url":
		return t.Url()
	default:
		return fmt.Errorf("unknown task type: %s", t.Type)
	}
}

// TODO what should happen when the user says no
func (t *Task) Confirm() error {
	err := confirm(t.Description)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) Cmd() error {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()

	args := strings.Split(t.Action, " ")
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
