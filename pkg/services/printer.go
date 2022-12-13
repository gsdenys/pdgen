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
package services

import (
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
)

// Printer interface that must be implemented by every printer.
type Printer interface {
	// SetWriter provides a way to overwrite the present writer. It's allow the user to
	//create a new output diferente of the default one, and redirect the output to the new one.
	SetWriter(path string) error

	// Init start the printer process
	Init(desc models.Describe)

	// Title function to print the title
	Title(title string)

	// SubSubtitle function to print the sub title
	Subtitle(subtitle string)

	// SubSubtitle function to print the sub sub title
	SubSubtitle(subSubtitle string)

	// LineBreak function to print a line break
	LineBreak()

	// Body function to print the text body
	Body(desc string)

	// Columns function that print the array of columns in a tabular format
	Columns(columns []models.Columns)

	// Table function to print a table documentation
	Table(t models.Table)

	// Done function that used to finalize the print process. In a file printed it can be used
	// to close the file.
	Done(desc models.Describe)
}

// PrintDocument is a default order to print each part of the Describe model.
func PrintDocument(p Printer, desc models.Describe) {
	p.Init(desc)
	t := translate.T

	p.Title(t.Sprintf("title"))
	p.LineBreak()

	p.Subtitle(t.Sprintf("title-db", desc.Database.Name))
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
