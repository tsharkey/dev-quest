/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"dev-quest/src/gamestate"
	"log"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the given quest log",
	Long:  `Use -l to specify the location of the quest log. If there isn't one specified then the default location will be used. This file is used as the base for the quests that the user will complete. It creates a new file in the users home directory if it doesn't exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("force").Value.String() == "true" {
			err := gamestate.Delete()
			if err != nil {
				log.Fatalf("Error deleting questing: %s", err)
			}
		}

		err := gamestate.InstallFrom(cmd.Flag("logfile").Value.String())
		if err != nil {
			log.Fatalf("Error installing quest log: %s", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().BoolP("force", "f", false, "Resets the user quest log")
}
