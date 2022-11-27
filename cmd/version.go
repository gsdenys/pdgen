/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print client version",
	Long:  `Print the pdgen client version information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pdgen version v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
