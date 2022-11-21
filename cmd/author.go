/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// authorCmd represents the author command
var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "Show the author informations",
	Long:  `Show all authors and contributors informations`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Author: Denys G. Santos")
	},
}

func init() {
	rootCmd.AddCommand(authorCmd)
}
