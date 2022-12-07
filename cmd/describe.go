/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/gsdenys/pdgen/pkg/options"
	"github.com/gsdenys/pdgen/pkg/services"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/spf13/cobra"
)

const (
	defaultDatabaseURI string = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

var (
	uri      string
	format   string
	schema   string
	database string
	path     string
	lang     string
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Generate the data dictionary output",
	Long:  "Connect to the database and generate the data dictionary output in the selected format that by default is a txt expressed at the standard output.",
	Run: func(cmd *cobra.Command, args []string) {
		//Output language definition
		if lang == "" {
			translate.InitLanguage()
		} else {
			if !translate.SetLanguage(lang) {
				fmt.Printf(
					"Sorry, the language %s is not registered, try to use another: %v\n",
					lang,
					translate.GetKeys(),
				)
				return
			}
		}

		//output format definition
		oFormat := options.Options[strings.ToUpper(format)]
		if oFormat == nil {
			fmt.Printf(
				"The format %s is not acceptable, please select one of: %v\n",
				string(format),
				options.GetKeys(),
			)
			return
		}

		//output path definition
		if path == "" {
			path = "output." + strings.ToLower(format)
		}

		//execute extractions
		desc, err := services.Describe(uri, database, schema)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//set path for output and execute print
		oFormat.SetWriter(path)
		services.PrintDocument(oFormat, *desc)

		if strings.ToUpper(format) != "DEFAULT" {
			fmt.Printf("%s document created.\n", strings.ToUpper(format))
		}
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)

	describeCmd.PersistentFlags().StringVarP(&format, "format", "f", "DEFAULT", options.Message())
	describeCmd.PersistentFlags().StringVarP(&lang, "language", "l", "", "The language selected to the output file")

	describeCmd.PersistentFlags().StringVarP(&uri, "uri", "u", defaultDatabaseURI, "The database connection uri")

	describeCmd.PersistentFlags().StringVarP(&database, "database", "d", "postgres", "The database to be described")
	describeCmd.PersistentFlags().StringVarP(&schema, "schema", "s", "public", "The schema to be described")

	describeCmd.PersistentFlags().StringVarP(&path, "output", "o", "", "The output file path")
}
