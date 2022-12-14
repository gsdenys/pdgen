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

const wantColumnsMd = `| table-title-name | table-title-type | table-title-allow | table-title-comment |
| :--- | :--- | :----: | :--- |
| name | text |  | Somme comment |
| type | text | BASE EXTENSION | Another comment |
`

func TestMD_SetWriter(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &MD{}
	assert.Nil(t, p.Out)

	_ = p.SetWriter(file)
	assert.NotNil(t, p.Out)
}

func TestMD_Init(t *testing.T) {
	p := &MD{}
	assert.Nil(t, p.Out)

	p.Init(models.Describe{})
	assert.Nil(t, p.Out)
}

func TestMD_Title(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &MD{Out: oFile}

	p.Title("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "# TEST\n")
}

func TestMD_Subtitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &MD{Out: oFile}

	p.Subtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "## TEST\n")
}

func TestMD_SubSubtitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &MD{Out: oFile}

	p.SubSubtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "### TEST\n")
}

func TestMD_Body(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &MD{Out: oFile}

	p.Body("Some test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "Some test\n")
}

func TestMD_LineBreak(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &MD{Out: oFile}

	p.LineBreak()

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\n")
}

func TestMD_Columns(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &MD{Out: oFile}

	translate.InitLanguage()
	p.Columns(columns)

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), wantColumnsMd)
}

func TestMD_Table(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &MD{Out: oFile}

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

	want := "### TEST\nTable test\n\n" + wantColumnsMd + "\n"

	assert.Equal(t, string(f), want)
}

func TestMD_Done(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &MD{Out: oFile}

	p.Done(models.Describe{})
}
