package services

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/rodaine/table"
)

type PrinterText struct {
	out      io.Writer
	formater bool
	language string
}

func (p *PrinterText) Title(title string) {
	if p.formater {
		fmt.Fprintf(p.out, "%s%s%s\n", string("\033[34m"), strings.ToUpper(title), string("\033[0m"))
	} else {
		fmt.Fprintf(p.out, "%s\n", strings.ToUpper(title))
	}
}

func (p *PrinterText) Subtitle(subtitle string) {
	p.Title(subtitle)
}

func (p *PrinterText) LineBreak() {
	fmt.Fprintf(p.out, "\n")
}

func (p *PrinterText) Body(desc string) {
	if p.formater {
		fmt.Fprintf(p.out, "%s%s\n", string("\033[0m"), desc)
	} else {
		fmt.Fprintf(p.out, "%s\n", desc)
	}
}

func (p *PrinterText) Columns(columns []models.Columns) {
	t := translate.GetTranslation(p.language)

	table.DefaultWriter = p.out
	tbl := table.New(
		t.Sprintf("table-title-name"),
		t.Sprintf("table-title-type"),
		t.Sprintf("table-title-allow"),
		t.Sprintf("table-title-comment"),
	)

	if p.formater {
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	}

	for c := range columns {
		tbl.AddRow(columns[c].Column, columns[c].Type, columns[c].Allow, columns[c].Comment)
	}

	tbl.Print()
}

func (p *PrinterText) Table(t models.Table) {
	p.Title(t.Name)
	p.Body(t.Desc)

	p.LineBreak()

	p.Columns(t.Columns)

	p.LineBreak()
}

func (p *PrinterText) print(desc models.Describe) {
	t := translate.GetTranslation(p.language)

	p.Title(t.Sprintf("title-db", desc.Database.Name))
	p.LineBreak()

	p.Body(desc.Database.Desc)
	p.LineBreak()

	// p.Title("Schema " + desc.Schema.Name)
	p.Title(t.Sprintf("title-schema", desc.Schema.Name))
	p.Body(desc.Schema.Desc)
	p.LineBreak()

	p.Title(t.Sprintf("title-tables"))
	p.Body(
		t.Sprintf("desc-tables",
			desc.Database.Name,
			desc.Schema.Name,
			len(desc.Tables),
		),
	)
	p.LineBreak()

	for index := range desc.Tables {
		p.Table(desc.Tables[index])
	}
}

func PrintDescription(desc models.Describe, lang string) {
	printer := &PrinterText{
		out:      os.Stdout,
		formater: true,
		language: lang,
	}

	printer.print(desc)
}
