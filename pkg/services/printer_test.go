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
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/uuid"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/gsdenys/pdgen/pkg/services/writer"
	"github.com/stretchr/testify/assert"
)

const fullWant = `TITLE

TITLE-DB
Database test

TITLE-SCHEMA
Schema test

TITLE-TABLES
desc-tables

TEST
Some description

table-title-name  table-title-type  table-title-allow  table-title-comment  
name              text                                 the name of test     
type              text              TEST PROD HML      the type of test     

`

func getWorkDir() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b) + "/test/"

	_ = os.MkdirAll(basepath, os.ModePerm)

	return basepath
}

func TestPrintDocument(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	desc := models.Describe{
		Database: models.Basic{
			Name: "test",
			Desc: "Database test",
		},
		Schema: models.Basic{
			Name: "test",
			Desc: "Schema test",
		},
		Tables: []models.Table{
			{
				Name: "test",
				Desc: "Some description",
				Columns: []models.Columns{
					{
						Column:  "name",
						Type:    "text",
						Allow:   "",
						Comment: "the name of test",
					},
					{
						Column:  "type",
						Type:    "text",
						Allow:   "TEST PROD HML",
						Comment: "the type of test",
					},
				},
			},
		},
	}

	translate.InitLanguage()

	oFile, err := writer.CreateFile(file)
	if err != nil {
		t.Error(err)
	}

	p := &writer.TXT{Out: oFile}

	PrintDocument(p, desc)

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), fullWant)
}
