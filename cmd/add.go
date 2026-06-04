/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Uptique/internal/app/add"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var url string
		if len(args) > 0 {
			url = args[0]
		}
		return add.Website(url)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
