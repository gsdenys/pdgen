package writer

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/rodaine/table"
	"golang.org/x/text/message"
)

type PrinterTXT struct {
	Out       io.Writer
	Translate *message.Printer
}

func (p *PrinterTXT) Init(desc models.Describe) {}

func (p *PrinterTXT) Title(title string) {
	fmt.Fprintf(p.Out, "%s\n", strings.ToUpper(title))
}

func (p *PrinterTXT) Subtitle(subtitle string) {
	p.Title(subtitle)
}

func (p *PrinterTXT) SubSubtitle(subtitle string) {
	p.Title(subtitle)
}

func (p *PrinterTXT) LineBreak() {
	fmt.Fprintf(p.Out, "\n")
}

func (p *PrinterTXT) Body(desc string) {
	fmt.Fprintf(p.Out, "%s\n", desc)
}

func (p *PrinterTXT) Columns(columns []models.Columns) {
	table.DefaultWriter = p.Out
	tbl := table.New(
		p.Translate.Sprintf("table-title-name"),
		p.Translate.Sprintf("table-title-type"),
		p.Translate.Sprintf("table-title-allow"),
		p.Translate.Sprintf("table-title-comment"),
	)

	for c := range columns {
		tbl.AddRow(columns[c].Column, columns[c].Type, columns[c].Allow, columns[c].Comment)
	}

	tbl.Print()
}

func (p *PrinterTXT) Table(t models.Table) {
	p.Title(t.Name)
	p.Body(t.Desc)

	p.LineBreak()

	p.Columns(t.Columns)

	p.LineBreak()
}

func (p *PrinterTXT) Done(desc models.Describe) {
	p.Out.(*os.File).Close()
}

func (p *PrinterTXT) GetLanguage() *message.Printer {
	return p.Translate
}
