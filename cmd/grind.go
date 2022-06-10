/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"dev-quest/src/quest"
	"log"

	"github.com/spf13/cobra"
)

var grindCmd = &cobra.Command{
	Use:   "grind",
	Short: "Grind all available quests",
	Long:  `Gets all available quests for the user to complete. You can choose one quest at a time to complete. Once a quest is completed the quest log will be updated and the user will be prompted to choose another quest. This happens until all quests are completed.`,
	Run: func(cmd *cobra.Command, args []string) {
		questLog, err := quest.GetQuestLog()
		if err != nil {
			log.Fatalf("error getting quest log: %v", err)
			return
		}

		err = questLog.Grind()
		if err != nil {
			log.Fatalf("error grinding quests: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(grindCmd)
}
