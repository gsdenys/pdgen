package writer

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gsdenys/pdgen/pkg/models"
	"golang.org/x/text/message"
)

const base string = `<!DOCTYPE html>
<html>

<head>
	<meta charset="UTF-8">
	<title>Page Title</title>
	<style>
		h1, h2, h3 {
			color: #009879;
			font-family: sans-serif;
		}
		.styled-table {
			border-collapse: collapse;
			margin: 25px 0;
			font-size: 0.9em;
			font-family: sans-serif;
			min-width: 400px;
			box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);
		}

		.styled-table thead tr {
			background-color: #009879;
			color: #ffffff;
			text-align: left;
		}

		.styled-table th,
		.styled-table td {
			padding: 12px 15px;
		}

		.styled-table tbody tr {
			border-bottom: 1px solid #dddddd;
		}
		
		.styled-table tbody tr:nth-of-type(even) {
			background-color: #f3f3f3;
		}
		
		.styled-table tbody tr:last-of-type {
			border-bottom: 2px solid #009879;
		}

		.styled-table tbody tr.active-row {
			font-weight: bold;
			color: #009879;
		}
	</style>
</head>

<body>`

type PrinterHTML struct {
	Out       io.Writer
	Translate *message.Printer
}

func (p *PrinterHTML) Init(desc models.Describe) {
	fmt.Fprintf(p.Out, "%s\n", base)
}

func (p *PrinterHTML) Title(title string) {
	fmt.Fprintf(p.Out, "\t<h1>%s</h1>\n", strings.ToUpper(title))
}

func (p *PrinterHTML) Subtitle(subtitle string) {
	fmt.Fprintf(p.Out, "\t<h2>%s</h2>\n", strings.ToUpper(subtitle))
}

func (p *PrinterHTML) SubSubtitle(subSubtitle string) {
	fmt.Fprintf(p.Out, "\t<h3>%s</h3>\n", strings.ToUpper(subSubtitle))
}

func (p *PrinterHTML) LineBreak() {
	fmt.Fprintf(p.Out, "\t<br>\n")
}

func (p *PrinterHTML) Body(desc string) {
	fmt.Fprintf(p.Out, "\t<p>%s</p>\n", desc)
}

func (p *PrinterHTML) Columns(columns []models.Columns) {
	fmt.Fprintf(p.Out, "\t%s\n", `<table class="styled-table">`)

	//Table title
	fmt.Fprintf(p.Out, "\t\t%s\n", `<tr class="active-row">`)
	fmt.Fprintf(p.Out, "\t\t\t<th>%s</th>\n", p.Translate.Sprintf("table-title-name"))
	fmt.Fprintf(p.Out, "\t\t\t<th>%s</th>\n", p.Translate.Sprintf("table-title-type"))
	fmt.Fprintf(p.Out, "\t\t\t<th>%s</th>\n", p.Translate.Sprintf("table-title-allow"))
	fmt.Fprintf(p.Out, "\t\t\t<th>%s</th>\n", p.Translate.Sprintf("table-title-comment"))
	fmt.Fprint(p.Out, "\t\t</tr>\n")

	for c := range columns {

		//Table title
		fmt.Fprint(p.Out, "\t\t<tr>\n")
		fmt.Fprintf(p.Out, "\t\t\t<td>%s</td>\n", columns[c].Column)
		fmt.Fprintf(p.Out, "\t\t\t<td>%s</td>\n", columns[c].Type)
		fmt.Fprintf(p.Out, "\t\t\t<td>%s</td>\n", columns[c].Allow)
		fmt.Fprintf(p.Out, "\t\t\t<td>%s</td>\n", columns[c].Comment)
		fmt.Fprint(p.Out, "\t\t</tr>\n")
	}

	fmt.Fprintf(p.Out, "\t</table>\n")
}

func (p *PrinterHTML) Table(t models.Table) {
	p.SubSubtitle(t.Name)
	p.Body(t.Desc)

	p.LineBreak()

	p.Columns(t.Columns)

	p.LineBreak()
}

func (p *PrinterHTML) Done(desc models.Describe) {
	fmt.Fprintf(p.Out, "\n</body>\n\n</html>")
	p.Out.(*os.File).Close()
}

func (p *PrinterHTML) GetLanguage() *message.Printer {
	return p.Translate
}
