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

		div {
			text-align: center;
		}

		svg g path {
			fill: #009879;
		}

		hr {
			height: 1px;
			background-color: #009879;
			border: none;
		}
	</style>
</head>

<body>`

const footer string = `
	<hr>
	<div>
		<svg version="1.0" xmlns="http://www.w3.org/2000/svg" width="10%" viewBox="0 0 500.000000 150.000000"
			preserveAspectRatio="xMidYMid meet">
			<g transform="translate(0.000000,150.000000) scale(0.100000,-0.100000)" fill="#000000" stroke="none">
				<path d="M567 1109 c-123 -49 -251 -169 -310 -288 -29 -60 -32 -75 -35 -176
	-3 -103 6 -165 26 -165 5 0 15 15 24 33 45 91 148 171 298 232 169 68 269 148
	323 259 32 65 29 74 -34 107 -54 27 -223 26 -292 -2z" />
				<path fill="#AB7C94" d="M2924 1046 c-84 -19 -139 -49 -195 -107 -58 -60 -79 -117 -79 -218 1
	-165 104 -275 285 -302 88 -13 201 0 271 31 l44 19 0 96 c0 108 6 127 50 145
	18 8 30 20 30 31 0 18 -10 19 -145 19 -136 0 -145 -1 -145 -19 0 -14 12 -23
	42 -33 31 -9 44 -20 49 -38 12 -43 10 -155 -3 -173 -19 -25 -108 -39 -168 -26
	-59 12 -130 71 -162 133 -33 64 -37 186 -8 254 74 172 292 206 385 60 20 -32
	37 -47 52 -47 21 -1 22 2 17 52 -9 94 -10 95 -65 111 -76 23 -185 28 -255 12z" />
				<path d="M1260 1031 c0 -14 11 -23 38 -31 62 -20 62 -22 62 -260 0 -240 -2
	-249 -65 -270 -19 -6 -35 -18 -35 -26 0 -11 28 -14 160 -14 135 0 160 2 160
	15 0 8 -12 17 -27 20 -16 3 -39 11 -53 18 l-25 13 -3 252 -2 252 52 0 c104 -1
	178 -57 178 -136 0 -87 -38 -122 -142 -131 -70 -7 -78 -27 -15 -44 96 -26 220
	18 258 91 25 45 25 125 0 170 -39 73 -152 99 -428 100 -105 0 -113 -1 -113
	-19z" />
				<path d="M1870 1034 c0 -14 12 -24 42 -34 64 -23 68 -36 68 -270 0 -232 -2
	-241 -70 -261 -22 -7 -40 -19 -40 -26 0 -12 38 -14 233 -11 212 3 236 5 278
	24 215 100 240 414 43 530 -74 44 -135 54 -351 61 -196 6 -203 5 -203 -13z
	m389 -50 c110 -32 174 -130 175 -264 1 -66 -3 -82 -30 -130 -43 -76 -100 -111
	-191 -117 -59 -5 -72 -2 -95 17 l-28 21 0 245 0 244 58 0 c32 0 82 -7 111 -16z" />
				<path d="M3360 1031 c0 -14 11 -23 38 -31 20 -7 43 -19 50 -28 13 -18 18 -457
	4 -478 -4 -7 -26 -17 -49 -23 -24 -7 -43 -18 -43 -25 0 -12 51 -15 278 -18
	152 -2 280 1 283 6 11 14 32 156 25 163 -12 12 -39 -11 -62 -52 -32 -58 -72
	-75 -175 -75 -127 0 -133 6 -137 149 l-4 111 73 -1 c97 0 138 -15 158 -56 9
	-18 24 -33 34 -33 15 0 17 12 17 105 0 97 -1 105 -19 105 -11 0 -22 -9 -26
	-20 -14 -44 -52 -60 -146 -60 l-89 0 0 115 0 115 109 0 c137 0 156 -7 181 -65
	14 -32 25 -45 40 -45 18 0 20 6 20 80 l0 80 -280 0 c-270 0 -280 -1 -280 -19z" />
				<path d="M4010 1031 c0 -13 11 -23 34 -31 19 -6 46 -20 60 -31 l26 -20 0 -213
	c0 -238 -2 -244 -69 -265 -24 -7 -41 -18 -41 -26 0 -12 23 -15 136 -15 122 0
	135 2 132 17 -2 10 -17 20 -38 24 -66 14 -70 29 -70 233 l1 181 144 -150 c79
	-82 177 -187 217 -232 64 -72 77 -82 107 -83 28 0 32 3 26 17 -10 28 -7 501 4
	521 5 10 28 27 50 37 27 12 41 25 41 37 0 17 -10 18 -131 18 -128 0 -130 0
	-127 -22 2 -15 13 -24 35 -30 60 -15 68 -41 71 -233 l3 -170 -221 228 -222
	227 -84 0 c-76 0 -84 -2 -84 -19z" />
				<path d="M907 972 c-68 -110 -178 -194 -322 -247 -136 -51 -246 -138 -291
	-233 -29 -61 -29 -61 50 -103 33 -18 53 -20 145 -17 90 2 115 7 163 30 120 56
	232 167 290 288 26 53 32 81 36 152 5 95 -1 145 -21 172 -12 16 -17 12 -50
	-42z" />
			</g>
		</svg>
	</div>`

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
	fmt.Fprintf(p.Out, "%s\n", footer)
	fmt.Fprintf(p.Out, "\n</body>\n\n</html>")

	_ = p.Out.(*os.File).Close()
}

func (p *PrinterHTML) GetLanguage() *message.Printer {
	return p.Translate
}
