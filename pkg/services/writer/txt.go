package writer

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/rodaine/table"
)

type TXT struct {
	Out io.Writer
}

func (p *TXT) SetWriter(path string) error {
	file, err := CreateFile(path)

	if err != nil {
		return err
	}

	p.Out = file
	return nil
}

func (p *TXT) Init(desc models.Describe) {
	// Do nothing because have nothing to initialise
}

func (p *TXT) Title(title string) {
	fmt.Fprintf(p.Out, "%s\n", strings.ToUpper(title))
}

func (p *TXT) Subtitle(subtitle string) {
	p.Title(subtitle)
}

func (p *TXT) SubSubtitle(subtitle string) {
	p.Title(subtitle)
}

func (p *TXT) LineBreak() {
	fmt.Fprintf(p.Out, "\n")
}

func (p *TXT) Body(desc string) {
	fmt.Fprintf(p.Out, "%s\n", desc)
}

func (p *TXT) Columns(columns []models.Columns) {
	table.DefaultWriter = p.Out
	tbl := table.New(
		translate.T.Sprintf("table-title-name"),
		translate.T.Sprintf("table-title-type"),
		translate.T.Sprintf("table-title-allow"),
		translate.T.Sprintf("table-title-comment"),
	)

	for c := range columns {
		tbl.AddRow(columns[c].Column, columns[c].Type, columns[c].Allow, columns[c].Comment)
	}

	tbl.Print()
}

func (p *TXT) Table(t models.Table) {
	p.Title(t.Name)
	p.Body(t.Desc)

	p.LineBreak()

	p.Columns(t.Columns)

	p.LineBreak()
}

func (p *TXT) Done(desc models.Describe) {
	_ = p.Out.(*os.File).Close()
}
