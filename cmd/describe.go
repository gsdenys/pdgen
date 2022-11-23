/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gsdenys/pdgen/pkg/options"
	"github.com/gsdenys/pdgen/pkg/services"
	"github.com/spf13/cobra"
)

var (
	uri      string                = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	format   options.OutputOptions = options.Options["DEFAULT"]
	schema   string                = "default"
	database string                = "postgres"
	path     string                = ""
)

const (
	defaultDatabaseURI   string = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	databaseUriName      string = "uri"
	databaseUriShorthand string = "u"
	databaseUriFlagDesc  string = "the database connection uri"

	defaultDatabase   string = "postgres"
	databaseName      string = "database"
	databaseShorthand string = "d"
	databaseFlagDesc  string = "the database to be described"

	defaultSchema   string = "public"
	schemaName      string = "schema"
	schemaShorthand string = "s"
	schemaFlagDesc  string = "the schema to be described"

	defaultPath   string = ""
	pathName      string = "out"
	pathShorthand string = "o"
	pathFlagDesc  string = "the description output file"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Generate the data dictionary output",
	Long: `Connect to the database and generate the data dictionary output 
in the selected format that by default is a txt expressed at the
standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		oFormat := options.OutputOptions(format)

		jsonStr, err := services.Describe(uri, database, schema)
		if err != nil {
			println(err.Error())
		}

		if path == "" {
			path = "output." + oFormat.String()
		}

		switch oFormat {
		case options.Options["JSON"]:
			services.ToJSON(*jsonStr, path)
		default:
			services.PrintDescription(*jsonStr)
		}
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)

	describeCmd.Flags().VarP(&format, "format", "f", options.Message())

	describeCmd.PersistentFlags().StringVarP(
		&uri,
		databaseUriName,
		databaseUriShorthand,
		defaultDatabaseURI,
		databaseUriFlagDesc,
	)

	describeCmd.PersistentFlags().StringVarP(
		&database,
		databaseName,
		databaseShorthand,
		defaultDatabase,
		databaseFlagDesc,
	)

	describeCmd.PersistentFlags().StringVarP(
		&schema,
		schemaName,
		schemaShorthand,
		defaultSchema,
		schemaFlagDesc,
	)

	describeCmd.PersistentFlags().StringVarP(
		&path,
		pathName,
		pathShorthand,
		defaultPath,
		pathFlagDesc,
	)
}
