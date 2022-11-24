package services

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
)

type PrinterMD struct {
	out      io.Writer
	language string
}

func (p *PrinterMD) Title(title string) {
	fmt.Fprintf(p.out, "# %s\n", strings.ToUpper(title))
}

func (p *PrinterMD) Subtitle(subtitle string) {
	fmt.Fprintf(p.out, "## %s\n", strings.ToUpper(subtitle))
}

func (p *PrinterMD) SubSubtitle(subtitle string) {
	fmt.Fprintf(p.out, "### %s\n", strings.ToUpper(subtitle))
}

func (p *PrinterMD) LineBreak() {
	fmt.Fprintf(p.out, "\n")
}

func (p *PrinterMD) Body(desc string) {
	fmt.Fprintf(p.out, "%s\n", desc)
}

func (p *PrinterMD) Columns(columns []models.Columns) {
	t := translate.GetTranslation(p.language)

	fmt.Fprintf(
		p.out,
		"| %s | %s | %s | %s |\n",
		t.Sprintf("table-title-name"),
		t.Sprintf("table-title-type"),
		t.Sprintf("table-title-allow"),
		t.Sprintf("table-title-comment"),
	)

	fmt.Fprintf(p.out, "| :--- | :--- | :----: | :--- |\n")

	for c := range columns {
		fmt.Fprintf(
			p.out,
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

func (p *PrinterMD) Print(desc models.Describe) {
	t := translate.GetTranslation(p.language)

	p.Title(t.Sprintf("title-db", desc.Database.Name))
	p.LineBreak()

	p.Body(desc.Database.Desc)
	p.LineBreak()

	p.Subtitle(t.Sprintf("title-schema", desc.Schema.Name))
	p.Body(desc.Schema.Desc)
	p.LineBreak()

	p.Subtitle(t.Sprintf("title-tables"))
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

func ToMD(desc models.Describe, path string, lang string) bool {
	var out bytes.Buffer
	printer := &PrinterMD{
		out:      &out,
		language: lang,
	}

	printer.Print(desc)

	err := ioutil.WriteFile(path, out.Bytes(), 0644)

	if err != nil {
		println("Error when try to save MD file!")
		println(out.String())
		return false
	}

	return true
}
