/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"dev-quest/src/game"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO: allows the user to export config to an md file
// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
		g := new(game.Game)
		err := viper.Unmarshal(g)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		log.Printf("%+v", g)

		by, err := ioutil.ReadFile("./templates/base.md")

		t, err := template.New("export").Funcs(sprig.FuncMap()).Parse(string(by))

		// write template to file
		f, err := os.Create("./export.md")
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		defer f.Close()

		err = t.Execute(f, g)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
