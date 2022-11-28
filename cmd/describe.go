/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"os"

	"github.com/gsdenys/pdgen/pkg/options"
	"github.com/gsdenys/pdgen/pkg/services"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/gsdenys/pdgen/pkg/services/writer"
	"github.com/spf13/cobra"
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

	defaultLang   string = "en"
	lanName       string = "language"
	langShorthand string = "l"
	langFlagDesc  string = "the language selected to the output file"
)

var (
	uri      string                = defaultDatabase
	format   options.OutputOptions = options.Options["DEFAULT"]
	schema   string                = defaultSchema
	database string                = defaultDatabase
	path     string                = ""
	lang     string                = "en"
)

func createFile(path string) io.Writer {
	file, err := os.Create(path)

	if err != nil {
		panic("Error create file: " + err.Error())
	}

	return file
}

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Generate the data dictionary output",
	Long: `Connect to the database and generate the data dictionary output 
in the selected format that by default is a txt expressed at the
standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		oFormat := options.OutputOptions(format)

		desc, err := services.Describe(uri, database, schema)
		if err != nil {
			println(err.Error())
		}

		if path == "" {
			path = "output." + oFormat.String()
		}

		translate.RegisterLanguages()
		translate := translate.GetTranslation(lang)

		var printer services.Printer
		switch oFormat {
		case options.Options["JSON"]:
			printer = &writer.PrinterJson{
				Out:       createFile(path),
				Translate: translate,
			}
		case options.Options["TXT"]:
			printer = &writer.PrinterTXT{
				Out:       createFile(path),
				Translate: translate,
			}
		case options.Options["MD"]:
			printer = &writer.PrinterMD{
				Out:       createFile(path),
				Translate: translate,
			}
		case options.Options["HTML"]:
			printer = &writer.PrinterHTML{
				Out:       createFile(path),
				Translate: translate,
			}
		default:
			printer = &writer.PrinterConsole{
				Out:       os.Stdout,
				Translate: translate,
			}
		}

		services.PrintDocument(printer, *desc)
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

	describeCmd.PersistentFlags().StringVarP(
		&lang,
		lanName,
		langShorthand,
		defaultLang,
		langFlagDesc,
	)
}
