/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"dev-quest/src/quest"
	"fmt"

	"github.com/spf13/cobra"
)

// grindCmd represents the grind command
var grindCmd = &cobra.Command{
	Use:   "grind",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grind called")

		questLog, err := quest.NewQuestLog(cmd.Flag("logfile").Value.String())
		if err != nil {
			fmt.Println(err)
		}

		err = questLog.Grind()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(grindCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grindCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grindCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
