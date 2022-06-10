package quest

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

func confirm(label string) error {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	res, err := prompt.Run()
	if err != nil {
		return err
	}

	if res != "y" {
		return fmt.Errorf("user did not confirm")
	}

	return nil
}

func completeQuest(questID string) error {
	viper.Set("quests."+questID+".completed", true)
	return viper.WriteConfig()
}

// TODO: this is kinda odd how you have to do this. You need
// to get the quest then update the array of values. I am not
// sure if this is the best way to do this. I couldn't figure out
// how to just edit the array element directly.
func completeTask(questID string, taskIndex int) error {
	tasks := Tasks{}
	err := viper.UnmarshalKey("quests."+questID+".tasks", &tasks)
	if err != nil {
		return err
	}
	return nil
}
