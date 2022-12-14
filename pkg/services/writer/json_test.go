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
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/uuid"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/message"
)

func getWorkDir() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b) + "/test/"

	_ = os.MkdirAll(basepath, os.ModePerm)

	return basepath
}

var baseTest models.Describe = models.Describe{
	Database: models.Basic{
		Name: "postgres",
		Desc: "default database",
	},
	Schema: models.Basic{
		Name: "public",
		Desc: "default database",
	},
	Tables: []models.Table{
		{
			Name: "test",
			Desc: "somme test",
			Columns: []models.Columns{
				{
					Column:  "test",
					Type:    "text",
					Allow:   "",
					Comment: "nothing",
				},
			},
		},
	},
}

func TestPrinterJson_Init(t *testing.T) {
	p := &JSON{
		Out: bytes.NewBuffer([]byte{}),
	}
	p.Init(baseTest)
}

func TestPrinterJson_Title(t *testing.T) {
	p := &JSON{
		Out: bytes.NewBuffer([]byte{}),
	}
	p.Title("test")
}

func TestPrinterJson_Subtitle(t *testing.T) {
	p := &JSON{
		Out: bytes.NewBuffer([]byte{}),
	}
	p.Subtitle("test")
}

func TestPrinterJson_SubSubtitle(t *testing.T) {
	p := &JSON{
		Out: bytes.NewBuffer([]byte{}),
	}
	p.SubSubtitle("test")
}

func TestPrinterJson_LineBreak(t *testing.T) {
	p := &JSON{
		Out: bytes.NewBuffer([]byte{}),
	}
	p.LineBreak()
}

func TestPrinterJson_Body(t *testing.T) {
	p := &JSON{
		Out: bytes.NewBuffer([]byte{}),
	}
	p.Body("test")
}

func TestPrinterJson_Columns(t *testing.T) {
	p := &JSON{
		Out: bytes.NewBuffer([]byte{}),
	}
	p.Columns([]models.Columns{})
}

func TestPrinterJson_Table(t *testing.T) {
	p := &JSON{
		Out: bytes.NewBuffer([]byte{}),
	}
	p.Table(models.Table{})
}

func TestPrinterJson_Done(t *testing.T) {
	type fields struct {
		Path      string
		Translate *message.Printer
	}
	type args struct {
		desc models.Describe
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "successful",
			fields: fields{
				Path: getWorkDir() + uuid.NewString(),
			},
			args: args{
				desc: baseTest,
			},
			want: "{\n    \"database\": {\n        \"name\": \"postgres\",\n        \"description\": \"default database\"\n    },\n    \"schema\": {\n        \"name\": \"public\",\n        \"description\": \"default database\"\n    },\n    \"tables\": [\n        {\n            \"name\": \"test\",\n            \"description\": \"somme test\",\n            \"columns\": [\n                {\n                    \"column\": \"test\",\n                    \"type\": \"text\",\n                    \"allow\": \"\",\n                    \"comment\": \"nothing\"\n                }\n            ]\n        }\n    ]\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := CreateFile(tt.fields.Path)
			if err != nil {
				t.Error(err)
			}

			p := &JSON{Out: file}
			p.Done(tt.args.desc)

			b, err := os.ReadFile(tt.fields.Path)
			if err != nil {
				t.Error(err.Error())
			}

			assert.Equal(t, string(b), tt.want)
		})
	}
}

func TestJSON_SetWriter(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &JSON{}
	assert.Nil(t, p.Out)

	_ = p.SetWriter(file)
	assert.NotNil(t, p.Out)
}
