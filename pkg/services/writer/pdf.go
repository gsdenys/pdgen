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
	"github.com/jung-kurt/gofpdf"
)

type PDF struct {
	Out io.Writer
	pdf *gofpdf.Fpdf
}

const (
	font  string  = "Arial"
	style string  = ""
	bold  string  = "B"
	size  float64 = 6
)

// Color is a RGB set of ints; for a nice picker
// see https://www.w3schools.com/colors/colors_picker.asp
type Color struct {
	Red, Green, Blue int
}

func (p *PDF) SetWriter(path string) error {
	file, err := CreateFile(path)

	if err != nil {
		return err
	}

	p.Out = file
	return nil
}

func (p *PDF) Init(desc models.Describe) {
	p.pdf = gofpdf.New("P", "mm", "A4", "")
	p.pdf.AddPage()

	p.pdf.SetFont(font, style, 10)
	p.pdf.SetTextColor(0, 0, 0)
}

func (p *PDF) Title(title string) {
	p.pdf.SetFont(font, bold, 18)
	p.pdf.Write(size, fmt.Sprintf("%s\n", strings.ToUpper(title)))
}

func (p *PDF) Subtitle(subtitle string) {
	p.pdf.SetFont(font, bold, 14)
	p.pdf.Write(size+3, fmt.Sprintf("%s\n", strings.ToUpper(subtitle)))
}

func (p *PDF) SubSubtitle(subSubtitle string) {
	p.pdf.SetFont(font, bold, 12)
	p.pdf.Write(size+3, fmt.Sprintf("%s\n", strings.ToUpper(subSubtitle)))
}

func (p *PDF) LineBreak() {
	p.pdf.Write(size+2, "\n")
}

func (p *PDF) Body(desc string) {
	p.pdf.SetFont(font, style, 10)
	p.pdf.Write(size, desc)
}

func (p *PDF) Columns(columns []models.Columns) {
	// fmt.Fprintf(
	// 	p.Out,
	// 	"| %s | %s | %s | %s |\n",
	// 	translate.T.Sprintf("table-title-name"),
	// 	translate.T.Sprintf("table-title-type"),
	// 	translate.T.Sprintf("table-title-allow"),
	// 	translate.T.Sprintf("table-title-comment"),
	// )

	// fmt.Fprintf(p.Out, "| :--- | :--- | :----: | :--- |\n")

	// for c := range columns {
	// 	fmt.Fprintf(
	// 		p.Out,
	// 		"| %s | %s | %s | %s |\n",
	// 		columns[c].Column,
	// 		columns[c].Type,
	// 		columns[c].Allow,
	// 		columns[c].Comment,
	// 	)
	// }
}

func (p *PDF) Table(t models.Table) {
	p.SubSubtitle(t.Name)
	p.Body(t.Desc)

	p.LineBreak()

	p.Columns(t.Columns)

	p.LineBreak()
}

func (p *PDF) Done(desc models.Describe) {
	p.pdf.Output(p.Out)
	_ = p.Out.(*os.File).Close()
}
