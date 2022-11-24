package services

import (
	"github.com/gsdenys/pdgen/pkg/models"
)

type Printer interface {
	LineBreak()
	Title(string)
	Subtitle(string)
	SubSubtitle(string)
	Body(string)
	Columns(columns []models.Columns)
	Table(table models.Table)
	Print(desc models.Describe)
}
