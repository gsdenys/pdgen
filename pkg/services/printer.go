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

// Printer interface that must be implemented by every printer
type Printer interface {
	SetWriter(path string) error
	Init(desc models.Describe)
	Title(title string)
	Subtitle(subtitle string)
	SubSubtitle(subSubtitle string)
	LineBreak()
	Body(desc string)
	Columns(columns []models.Columns)
	Table(t models.Table)
	Done(desc models.Describe)
}

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
