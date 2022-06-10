package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

const (
	FileName   = ".dev-quest"
	FileType   = "yml"
	ConfigPath = "$HOME/"

	FullConfigPath = ConfigPath + FileName + "." + FileType
)

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

	err = setQuestIDs()
	if err != nil {
		return err
	}

	return viper.WriteConfigAs(absPathify(FullConfigPath))
}

func setQuestIDs() error {
	quests := viper.Get("quests").(map[string]interface{})
	for questID := range quests {
		viper.Set("quests."+questID+".id", questID)
	}

	return viper.WriteConfig()
}

func Delete() error {
	err := os.Remove(FullConfigPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}

func absPathify(inPath string) string {
	if inPath == "$HOME" || strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator)) {
		inPath = userHomeDir() + inPath[5:]
	}

	inPath = os.ExpandEnv(inPath)

	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}

	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}

	return ""
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
