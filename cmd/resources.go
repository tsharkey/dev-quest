/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"dev-quest/src/game"
	"log"

	"github.com/spf13/cobra"
)

// resourcesCmd represents the resources command
var resourcesCmd = &cobra.Command{
	Use:   "resources",
	Short: "List all the resources for the user",
	Long:  `A collection of urls for the user to reference at any point in time`,
	Run: func(cmd *cobra.Command, args []string) {
		resources, err := game.LoadResources()
		if err != nil {
			log.Fatalf("Error getting resources: %s", err)
		}

		err = resources.Display()
		if err != nil {
			log.Fatalf("Error displaying resources: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(resourcesCmd)
}
