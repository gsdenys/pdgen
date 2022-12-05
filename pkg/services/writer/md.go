package writer

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
)

type PrinterMD struct {
	Out io.Writer
}

func (p *PrinterMD) Init(desc models.Describe) {
	// Do nothing because have nothing to initialise
}

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
		translate.T.Sprintf("table-title-name"),
		translate.T.Sprintf("table-title-type"),
		translate.T.Sprintf("table-title-allow"),
		translate.T.Sprintf("table-title-comment"),
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
	_ = p.Out.(*os.File).Close()
}
