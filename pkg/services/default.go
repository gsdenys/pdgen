package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/rodaine/table"
)

type Printer struct{}

const (
	reset = "\033[0m"
	blue  = "\033[0m"
)

const lineBreak = "\n\n"

func printTitle(title string) {
	fmt.Fprintf(os.Stdout, "%s%s%s\n", string("\033[34m"), strings.ToUpper(title), string("\033[0m"))
}

func breakLine() {
	fmt.Fprintf(os.Stdout, "\n")
}

func printDescription(desc string) {
	fmt.Fprintf(os.Stdout, "%s%s\n", string("\033[0m"), desc)
}

func printTable(t models.Table) {
	printTitle(t.Name)
	printDescription(t.Desc)

	breakLine()

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Type", "Allow", "Comment")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for c := range t.Columns {
		tbl.AddRow(t.Columns[c].Column, t.Columns[c].Type, t.Columns[c].Allow, t.Columns[c].Comment)
	}

	tbl.Print()
	breakLine()

}

func PrintDescription(desc models.Describe) {
	printTitle("Data Dictionary for database " + desc.Database.Name)
	breakLine()

	printDescription(desc.Database.Desc)
	breakLine()

	printTitle("Schema " + desc.Schema.Name)
	printDescription(desc.Schema.Desc)
	breakLine()

	printTitle("Tables Descriptions")
	printDescription(
		fmt.Sprintf(
			`The database %s, at the schema %s, contem %d tables that are described bellow. 
For each table is presented their name, description and the description of each
column containing their name, type, and description. In the cases of the data 
type is a custom dada type, the options is printed as an allow  enum`,
			desc.Database.Name,
			desc.Schema.Name,
			len(desc.Tables),
		),
	)
	breakLine()

	for index := range desc.Tables {
		printTable(desc.Tables[index])
	}
}
