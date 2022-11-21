/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gsdenys/pdgen/pkg/options"
	"github.com/spf13/cobra"
)

var (
	uri string                = "postgres://postgres:postgres@localhost:5432/postgres"
	out options.OutputOptions = options.Options["TXT"]
)

const (
	defaultDatabaseURI   string = "postgres://postgres:postgres@localhost:5432/postgres"
	databaseUriName      string = "uri"
	databaseUriShorthand string = "u"
	databaseUriFlagDesc  string = "the database connection uri"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Generate the data dictionary output",
	Long: `Connect to the database and generate the data dictionary output 
in the selected format that by default is a txt expressed at the
standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("describe called", uri, out)
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)

	describeCmd.Flags().VarP(&out, "format", "f", options.Message())

	describeCmd.PersistentFlags().StringVarP(
		&uri,
		databaseUriName,
		databaseUriShorthand,
		defaultDatabaseURI,
		databaseUriFlagDesc,
	)
}
