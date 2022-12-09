package writer

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

type DEFAULT struct {
	Out io.Writer
}

func (p *DEFAULT) SetWriter(path string) error {
	p.Out = os.Stdout

	return nil
}

func (p *DEFAULT) Init(desc models.Describe) {
	p.Out = os.Stdout
}

func (p *DEFAULT) Title(title string) {
	fmt.Fprintf(p.Out, "%s%s%s\n", string("\033[0;32m"), strings.ToUpper(title), string("\033[0m"))
}

func (p *DEFAULT) Subtitle(subtitle string) {
	p.Title(subtitle)
}

func (p *DEFAULT) SubSubtitle(subtitle string) {
	p.Title(subtitle)
}

func (p *DEFAULT) LineBreak() {
	fmt.Fprintf(p.Out, "\n")
}

func (p *DEFAULT) Body(desc string) {
	fmt.Fprintf(p.Out, "%s%s\n", string("\033[0m"), desc)
}

func (p *DEFAULT) Columns(columns []models.Columns) {
	table.DefaultWriter = p.Out
	tbl := table.New(
		translate.T.Sprintf("table-title-name"),
		translate.T.Sprintf("table-title-type"),
		translate.T.Sprintf("table-title-allow"),
		translate.T.Sprintf("table-title-comment"),
	)

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for c := range columns {
		tbl.AddRow(columns[c].Column, columns[c].Type, columns[c].Allow, columns[c].Comment)
	}

	tbl.Print()
}

func (p *DEFAULT) Table(t models.Table) {
	p.Title(t.Name)
	p.Body(t.Desc)

	p.LineBreak()

	p.Columns(t.Columns)

	p.LineBreak()
}

func (p *DEFAULT) Done(desc models.Describe) {}
