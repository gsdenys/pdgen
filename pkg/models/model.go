package models

import (
	"strings"

	"bytes"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

const lineAfter = "\n\n\n"
const sepLine = "\n\n"

type Basic struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

func (b *Basic) String() string {
	var out string = ""

	out += strings.ToUpper(b.Name)
	out += sepLine
	out += b.Desc
	out += lineAfter

	return out
}

type Columns struct {
	Column  string `json:"column"`
	Type    string `json:"type"`
	Allow   string `json:"allow"`
	Comment string `json:"comment"`
}

type Table struct {
	Name    string    `json:"name"`
	Desc    string    `json:"description"`
	Columns []Columns `json:"columns"`
}

func (t *Table) String() string {
	var out string

	out += strings.ToUpper(t.Name)
	out += sepLine
	out += t.Desc
	out += sepLine

	var buf bytes.Buffer

	table.DefaultWriter = &buf
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Type", "Allow", "Comment")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for c := range t.Columns {
		tbl.AddRow(t.Columns[c].Column, t.Columns[c].Type, t.Columns[c].Allow, t.Columns[c].Comment)
	}

	tbl.Print()

	out += buf.String()
	out += sepLine

	return out
}

type Describe struct {
	Database Basic   `json:"database"`
	Schema   Basic   `json:"schema"`
	Tables   []Table `json:"tables"`
}

func (d *Describe) String() string {
	var out string = ""

	out += "Database: " + d.Database.String()
	out += "Schema: " + d.Schema.String()

	out += "TABLES"
	out += sepLine

	for index := range d.Tables {
		out += d.Tables[index].String()
	}

	return out
}
