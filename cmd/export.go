/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"dev-quest/src/export"
	"dev-quest/src/game"
	"html/template"
	"log"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
		g := new(game.Game)
		err := viper.Unmarshal(g)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		exportTemplate := export.Markdown
		if cmd.Flag("template").Value.String() != "" {
			t, err := getTemplateString(cmd.Flag("template").Value.String())
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			exportTemplate = t
		}

		t, err := template.New("export").Funcs(sprig.FuncMap()).Parse(exportTemplate)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// wri template to file
		f, err := os.Create(cmd.Flag("output").Value.String())
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
	exportCmd.Flags().StringP("output", "o", "export.md", "output file")
	exportCmd.Flags().StringP("template", "t", "", "template file")
}

func getTemplateString(templateFile string) (string, error) {
	by, err := os.ReadFile(templateFile)
	if err != nil {
		return "", err
	}

	return string(by), nil
}
