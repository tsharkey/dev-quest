/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"dev-quest/src/gamestate"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "dev-quest",
	Short: "An onboarding tool for developers",
	Long:  `Schema based onboarding. Create a single file for the user to use and have a list of quests for them to complete as part of their onboarding process.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("logfile", "l", "./quest.yml", "the yaml file to build the questlog from")

	viper.SetConfigType(gamestate.FileType)
	viper.SetConfigName(gamestate.FileName)
	viper.AddConfigPath(gamestate.ConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config... have you installed your quest file yet? %s", err)
	}
}
