package game

import (
	"bytes"
	"dev-quest/src/quest"
	"dev-quest/src/resource"
	"dev-quest/src/util"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

const (
	FileName       = ".dev-quest"
	FileType       = "yml"
	ConfigPath     = "$HOME/"
	FullConfigPath = ConfigPath + FileName + "." + FileType

	ConfigKeyResources  = "resources"
	ConfigKeyUserConfig = "user_config"
	ConfigKeyStorylines = "storylines"
)

func SaveStorylines(sls quest.Storylines) error {
	viper.Set(ConfigKeyStorylines, sls)
	return viper.WriteConfig()
}

func LoadStorylines() (quest.Storylines, error) {
	sls := new(quest.Storylines)
	return *sls, viper.UnmarshalKey(ConfigKeyStorylines, sls)
}

func LoadResources() (resource.Resources, error) {
	resources := new(resource.Resources)
	return *resources, viper.UnmarshalKey(ConfigKeyResources, resources)
}

func InstallFrom(baseConfigPath string) error {
	_, err := os.Stat(baseConfigPath)
	if err != nil {
		return err
	}

	// get file contents
	fileContents, err := ioutil.ReadFile(baseConfigPath)
	if err != nil {
		return err
	}

	err = viper.ReadConfig(bytes.NewBuffer(fileContents))
	if err != nil {
		return err
	}

	return viper.WriteConfigAs(util.GetAbsolutePath(FullConfigPath))
}

func Delete() error {
	err := os.Remove(FullConfigPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}
