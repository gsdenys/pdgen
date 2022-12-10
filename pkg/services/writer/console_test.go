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
	"bytes"
	"os"
	"testing"

	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/stretchr/testify/assert"
)

const wantColumnsConsole = `table-title-name  table-title-type  table-title-allow  table-title-comment  
name              text                                 Somme comment        
type              text              BASE EXTENSION     Another comment      
`

func TestDEFAULT_SetWriter(t *testing.T) {
	p := &DEFAULT{}
	assert.Nil(t, p.Out)

	_ = p.SetWriter("some/path")
	assert.NotNil(t, p.Out)
	assert.Equal(t, p.Out, os.Stdout)
}

func TestDEFAULT_Init(t *testing.T) {
	p := &DEFAULT{}
	assert.Nil(t, p.Out)

	p.Init(models.Describe{})
	assert.NotNil(t, p.Out)
	assert.Equal(t, p.Out, os.Stdout)
}

func TestDEFAULT_Title(t *testing.T) {
	var b []byte
	buf := bytes.NewBuffer(b)

	p := &DEFAULT{
		Out: buf,
	}

	p.Title("test")
	assert.Equal(t, buf.String(), "\x1b[0;32mTEST\x1b[0m\n")
}

func TestDEFAULT_Subtitle(t *testing.T) {
	var b []byte
	buf := bytes.NewBuffer(b)

	p := &DEFAULT{
		Out: buf,
	}

	p.Subtitle("test")
	assert.Equal(t, buf.String(), "\x1b[0;32mTEST\x1b[0m\n")
}

func TestDEFAULT_SubSubtitle(t *testing.T) {
	var b []byte
	buf := bytes.NewBuffer(b)

	p := &DEFAULT{
		Out: buf,
	}

	p.SubSubtitle("test")
	assert.Equal(t, buf.String(), "\x1b[0;32mTEST\x1b[0m\n")
}

func TestDEFAULT_LineBreak(t *testing.T) {
	var b []byte
	buf := bytes.NewBuffer(b)

	p := &DEFAULT{
		Out: buf,
	}

	p.LineBreak()
	assert.Equal(t, buf.String(), "\n")
}

func TestDEFAULT_Body(t *testing.T) {
	var b []byte
	buf := bytes.NewBuffer(b)

	p := &DEFAULT{
		Out: buf,
	}

	p.Body("test")
	assert.Equal(t, buf.String(), "\033[0mtest\n")
}

func TestDEFAULT_Columns(t *testing.T) {
	var b []byte
	buf := bytes.NewBuffer(b)

	p := &DEFAULT{
		Out: buf,
	}

	//columns is defined at html_test.go
	translate.InitLanguage()
	p.Columns(columns)

	assert.Equal(t, buf.String(), wantColumnsConsole)
}

func TestDEFAULT_Table(t *testing.T) {
	var b []byte
	buf := bytes.NewBuffer(b)

	p := &DEFAULT{
		Out: buf,
	}

	table := models.Table{
		Name:    "test",
		Desc:    "Table test",
		Columns: columns,
	}

	//columns is defined at html_test.go
	translate.InitLanguage()
	p.Table(table)

	want := "\x1b[0;32mTEST\x1b[0m\n\x1b[0mTable test\n\n" + wantColumnsConsole + "\n"

	assert.Equal(t, buf.String(), want)
}

func TestDEFAULT_Done(t *testing.T) {
	var b []byte
	buf := bytes.NewBuffer(b)

	p := &DEFAULT{
		Out: buf,
	}

	p.Done(models.Describe{})
	assert.Equal(t, buf.String(), "")
}
