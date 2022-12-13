/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package writer

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
)

const Base string = `<!DOCTYPE html>
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

const cell_string_pattern string = "\t\t\t<td>%s</td>\n"
const cell_string_header string = "\t\t\t<th>%s</th>\n"

type HTML struct {
	Out io.Writer
}

func (p *HTML) SetWriter(path string) error {
	file, err := CreateFile(path)

	if err != nil {
		return err
	}

	p.Out = file
	return nil
}

func (p *HTML) Init(desc models.Describe) {
	fmt.Fprintf(p.Out, "%s\n", Base)
}

func (p *HTML) Title(title string) {
	fmt.Fprintf(p.Out, "\t<h1>%s</h1>\n", strings.ToUpper(title))
}

func (p *HTML) Subtitle(subtitle string) {
	fmt.Fprintf(p.Out, "\t<h2>%s</h2>\n", strings.ToUpper(subtitle))
}

func (p *HTML) SubSubtitle(subSubtitle string) {
	fmt.Fprintf(p.Out, "\t<h3>%s</h3>\n", strings.ToUpper(subSubtitle))
}

func (p *HTML) LineBreak() {
	fmt.Fprintf(p.Out, "\t<br>\n")
}

func (p *HTML) Body(desc string) {
	fmt.Fprintf(p.Out, "\t<p>%s</p>\n", desc)
}

func (p *HTML) Columns(columns []models.Columns) {
	fmt.Fprintf(p.Out, "\t%s\n", `<table class="styled-table">`)

	//Table title
	fmt.Fprintf(p.Out, "\t\t%s\n", `<tr class="active-row">`)
	fmt.Fprintf(p.Out, cell_string_header, translate.T.Sprintf("table-title-name"))
	fmt.Fprintf(p.Out, cell_string_header, translate.T.Sprintf("table-title-type"))
	fmt.Fprintf(p.Out, cell_string_header, translate.T.Sprintf("table-title-allow"))
	fmt.Fprintf(p.Out, cell_string_header, translate.T.Sprintf("table-title-comment"))
	fmt.Fprint(p.Out, "\t\t</tr>\n")

	for c := range columns {

		//Table title
		fmt.Fprint(p.Out, "\t\t<tr>\n")

		fmt.Fprintf(p.Out, cell_string_pattern, columns[c].Column)
		fmt.Fprintf(p.Out, cell_string_pattern, columns[c].Type)
		fmt.Fprintf(p.Out, cell_string_pattern, columns[c].Allow)
		fmt.Fprintf(p.Out, cell_string_pattern, columns[c].Comment)
		fmt.Fprint(p.Out, "\t\t</tr>\n")
	}

	fmt.Fprintf(p.Out, "\t</table>\n")
}

func (p *HTML) Table(t models.Table) {
	p.SubSubtitle(t.Name)
	p.Body(t.Desc)

	p.LineBreak()

	p.Columns(t.Columns)

	p.LineBreak()
}

func (p *HTML) Done(desc models.Describe) {
	fmt.Fprintf(p.Out, "\n</body>\n\n</html>")

	_ = p.Out.(*os.File).Close()
}
