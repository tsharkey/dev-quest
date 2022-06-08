package quest

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/pkg/browser"
	"gopkg.in/yaml.v3"
)

type QuestLog struct {
	logFile   string     `yaml:"-"`
	Inits     []Init     `yaml:"inits"`
	Resources []Resource `yaml:"resources"`
	Installs  []Install  `yaml:"installs"`
	Clones    []Clone    `yaml:"clones"`
}

type Init struct {
	Name string
}

type Resource struct {
	Name string `yaml:"name"`
	Hint string `yaml:"hint"`
	URL  string `yaml:"url"`
}

type Install struct {
	Name   string
	Action string
}

func (i *Install) Install() error {
	prompt := promptui.Prompt{
		Label:     color.YellowString("do you need to install %s", i.Name),
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	if strings.ToLower(result) != "y" {
		return nil
	}

	args := strings.Split(i.Action, " ")
	cmd := exec.Command(args[0], args[1:]...)

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner                               // Run for some time to simulate work

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Install failed: %s", out)
		return err
	}

	s.Stop()

	color.Green("%s installed", i.Name)

	return nil
}

type Clone struct {
	Name string
}

func NewQuestLog(logFile string) (*QuestLog, error) {
	var ql QuestLog
	ql.logFile = logFile
	err := ql.Load(logFile)
	if err != nil {
		return nil, err
	}
	return &ql, nil
}

func (ql *QuestLog) Load(logFile string) error {
	ql.logFile = logFile

	fileContents, err := ioutil.ReadFile(logFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileContents, &ql)
	if err != nil {
		return err
	}

	return nil
}

func (ql *QuestLog) Grind() error {
	if err := ql.Install(); err != nil {
		return err
	}

	return nil
}

func (ql *QuestLog) Install() error {
	for _, install := range ql.Installs {
		if err := install.Install(); err != nil {
			return err
		}
	}

	return nil
}

func (ql *QuestLog) ShowResources() error {
	prompt := promptui.Select{
		Label: "Select a resource",
		Items: ql.getResourceDisplayStrings(),
	}

	ix, result, err := prompt.Run()

	log.Printf("result: %s", result)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	return browser.OpenURL(ql.Resources[ix].URL)
}

func (ql *QuestLog) getResourceDisplayStrings() []string {
	var strings []string
	for _, resource := range ql.Resources {
		strings = append(strings, resource.Name)
	}
	return strings
}
