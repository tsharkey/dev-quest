/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"dev-quest/src/quest"
	"fmt"

	"github.com/spf13/cobra"
)

// resourcesCmd represents the resources command
var resourcesCmd = &cobra.Command{
	Use:   "resources",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("resources called")

		questLog, err := quest.NewQuestLog(cmd.Flag("logfile").Value.String())
		if err != nil {
			fmt.Println(err)
		}

		if err := questLog.ShowResources(); err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(resourcesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resourcesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resourcesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
