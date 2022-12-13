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

func TestTXT_SetWriter(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &TXT{}
	assert.Nil(t, p.Out)

	_ = p.SetWriter(file)
	assert.NotNil(t, p.Out)
}

func TestTXT_Init(t *testing.T) {
	p := &TXT{}
	assert.Nil(t, p.Out)

	p.Init(models.Describe{})
	assert.Nil(t, p.Out)
}

func TestTXT_Title(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &TXT{Out: oFile}

	p.Title("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "TEST\n")
}

func TestTXT_Subtitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &TXT{Out: oFile}

	p.Subtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "TEST\n")
}

func TestTXT_SubSubTitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &TXT{Out: oFile}

	p.SubSubtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "TEST\n")
}

func TestTXT_LineBreak(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &TXT{Out: oFile}

	p.LineBreak()

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\n")
}

func TestTXT_Body(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &TXT{Out: oFile}

	p.Body("Some test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "Some test\n")
}

func TestTXT_Columns(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &TXT{Out: oFile}

	//columns is defined at html_test.go
	translate.InitLanguage()
	p.Columns(columns)

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), wantColumnsConsole)
}

func TestTXT_Table(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &TXT{Out: oFile}

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

	want := "TEST\nTable test\n\n" + wantColumnsConsole + "\n"

	assert.Equal(t, string(f), want)
}

func TestTXT_Done(t *testing.T) {
	file := getWorkDir() + uuid.NewString()
	oFile, err := CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &TXT{Out: oFile}

	p.Done(models.Describe{})
}
