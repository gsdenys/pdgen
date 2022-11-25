package services

import (
	"github.com/gsdenys/pdgen/pkg/models"
	"golang.org/x/text/message"
)

type Printer interface {
	Init(desc models.Describe)
	Title(title string)
	Subtitle(subtitle string)
	SubSubtitle(subSubtitle string)
	LineBreak()
	Body(desc string)
	Columns(columns []models.Columns)
	Table(t models.Table)
	Done(desc models.Describe)
	GetLanguage() *message.Printer
}

func PrintDocument(p Printer, desc models.Describe) {
	p.Init(desc)
	t := p.GetLanguage()

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

	p.Done(desc)
}
