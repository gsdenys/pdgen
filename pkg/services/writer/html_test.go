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
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/stretchr/testify/assert"
)

const wantColumnsHtml string = "\t<table class=\"styled-table\">\n\t\t<tr class=\"active-row\">\n\t\t\t<th>table-title-name</th>\n\t\t\t<th>table-title-type</th>\n\t\t\t<th>table-title-allow</th>\n\t\t\t<th>table-title-comment</th>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>name</td>\n\t\t\t<td>text</td>\n\t\t\t<td></td>\n\t\t\t<td>Somme comment</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>type</td>\n\t\t\t<td>text</td>\n\t\t\t<td>BASE EXTENSION</td>\n\t\t\t<td>Another comment</td>\n\t\t</tr>\n\t</table>\n"

var columns []models.Columns = []models.Columns{
	{
		Column:  "name",
		Type:    "text",
		Allow:   "",
		Comment: "Somme comment",
	},
	{
		Column:  "type",
		Type:    "text",
		Allow:   "BASE EXTENSION",
		Comment: "Another comment",
	},
}

func TestPrinterHTML_Init(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &HTML{Out: oFile}

	p.Init(models.Describe{})

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), Base+"\n")
}

func TestPrinterHTML_Title(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &HTML{Out: oFile}
	p.Title("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<h1>TEST</h1>\n")
}

func TestPrinterHTML_Subtitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &HTML{Out: oFile}

	p.Subtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<h2>TEST</h2>\n")
}

func TestPrinterHTML_SubSubtitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &HTML{Out: oFile}

	p.SubSubtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<h3>TEST</h3>\n")
}

func TestPrinterHTML_LineBreak(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &HTML{Out: oFile}

	p.LineBreak()

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<br>\n")
}

func TestPrinterHTML_Body(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &HTML{Out: oFile}

	p.Body("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<p>test</p>\n")
}

func TestHTML_Columns(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &HTML{Out: oFile}

	translate.InitLanguage()
	p.Columns(columns)

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), wantColumnsHtml)
}

func TestHTML_SetWriter(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &HTML{}
	assert.Nil(t, p.Out)

	_ = p.SetWriter(file)
	assert.NotNil(t, p.Out)
}

func TestHTML_Table(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &HTML{Out: oFile}

	translate.InitLanguage()

	table := models.Table{
		Name:    "test",
		Desc:    "Table test",
		Columns: columns,
	}

	p.Table(table)

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	want := "\t<h3>TEST</h3>\n\t<p>Table test</p>\n\t<br>\n" + wantColumnsHtml + "\t<br>\n"

	assert.Equal(t, string(f), want)
}

func TestHTML_Done(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &HTML{Out: oFile}

	translate.InitLanguage()
	p.Done(models.Describe{})

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\n</body>\n\n</html>")
}
