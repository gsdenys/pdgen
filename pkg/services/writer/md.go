package writer

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gsdenys/pdgen/pkg/models"
	"golang.org/x/text/message"
)

type PrinterMD struct {
	Out       io.Writer
	Translate *message.Printer
}

func (p *PrinterMD) Init(desc models.Describe) {}

func (p *PrinterMD) Title(title string) {
	fmt.Fprintf(p.Out, "# %s\n", strings.ToUpper(title))
}

func (p *PrinterMD) Subtitle(subtitle string) {
	fmt.Fprintf(p.Out, "## %s\n", strings.ToUpper(subtitle))
}

func (p *PrinterMD) SubSubtitle(subSubtitle string) {
	fmt.Fprintf(p.Out, "### %s\n", strings.ToUpper(subSubtitle))
}

func (p *PrinterMD) LineBreak() {
	fmt.Fprintf(p.Out, "\n")
}

func (p *PrinterMD) Body(desc string) {
	fmt.Fprintf(p.Out, "%s\n", desc)
}

func (p *PrinterMD) Columns(columns []models.Columns) {
	fmt.Fprintf(
		p.Out,
		"| %s | %s | %s | %s |\n",
		p.Translate.Sprintf("table-title-name"),
		p.Translate.Sprintf("table-title-type"),
		p.Translate.Sprintf("table-title-allow"),
		p.Translate.Sprintf("table-title-comment"),
	)

	fmt.Fprintf(p.Out, "| :--- | :--- | :----: | :--- |\n")

	for c := range columns {
		fmt.Fprintf(
			p.Out,
			"| %s | %s | %s | %s |\n",
			columns[c].Column,
			columns[c].Type,
			columns[c].Allow,
			columns[c].Comment,
		)
	}
}

func (p *PrinterMD) Table(t models.Table) {
	p.SubSubtitle(t.Name)
	p.Body(t.Desc)

	p.LineBreak()

	p.Columns(t.Columns)

	p.LineBreak()
}

func (p *PrinterMD) Done(desc models.Describe) {
	p.Out.(*os.File).Close()
}

func (p *PrinterMD) GetLanguage() *message.Printer {
	return p.Translate
}
