package writer

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/rodaine/table"
	"golang.org/x/text/message"
)

type PrinterConsole struct {
	Out       io.Writer
	Translate *message.Printer
}

func (p *PrinterConsole) Init(desc models.Describe) {}

func (p *PrinterConsole) Title(title string) {
	fmt.Fprintf(p.Out, "%s%s%s\n", string("\033[34m"), strings.ToUpper(title), string("\033[0m"))
}

func (p *PrinterConsole) Subtitle(subtitle string) {
	p.Title(subtitle)
}

func (p *PrinterConsole) SubSubtitle(subtitle string) {
	p.Title(subtitle)
}

func (p *PrinterConsole) LineBreak() {
	fmt.Fprintf(p.Out, "\n")
}

func (p *PrinterConsole) Body(desc string) {
	fmt.Fprintf(p.Out, "%s%s\n", string("\033[0m"), desc)
}

func (p *PrinterConsole) Columns(columns []models.Columns) {
	table.DefaultWriter = p.Out
	tbl := table.New(
		p.Translate.Sprintf("table-title-name"),
		p.Translate.Sprintf("table-title-type"),
		p.Translate.Sprintf("table-title-allow"),
		p.Translate.Sprintf("table-title-comment"),
	)

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for c := range columns {
		tbl.AddRow(columns[c].Column, columns[c].Type, columns[c].Allow, columns[c].Comment)
	}

	tbl.Print()
}

func (p *PrinterConsole) Table(t models.Table) {
	p.Title(t.Name)
	p.Body(t.Desc)

	p.LineBreak()

	p.Columns(t.Columns)

	p.LineBreak()
}

func (p *PrinterConsole) GetLanguage() *message.Printer {
	return p.Translate
}

func (p *PrinterConsole) Done(desc models.Describe) {}
