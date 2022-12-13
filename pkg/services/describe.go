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
	"github.com/gsdenys/pdgen/pkg/database"
	"github.com/gsdenys/pdgen/pkg/models"
)

// Describe function that has the main objective create the describe data structure that
// represents all elements of database at provided schema
func Describe(uri string, db string, schema string) (*models.Describe, error) {
	desc := &models.Describe{}

	conn, err := database.Connect("postgres", uri)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// get Database Info
	dbDesc, err := database.GetDatabaseComment(conn, db)
	if err != nil {
		return nil, err
	}

	desc.Database = models.Basic{
		Name: db,
		Desc: dbDesc,
	}

	//get schema info
	scDesc, err := database.GetSchemaComment(conn, schema)
	if err != nil {
		return nil, err
	}

	desc.Schema = models.Basic{
		Name: schema,
		Desc: scDesc,
	}

	tables, err := database.GetAllTables(conn, schema)
	if err != nil {
		return nil, err
	}

	desc.Tables = tables

	for i := range desc.Tables {
		columns, err := database.GetTableColumns(conn, schema, desc.Tables[i].Name)
		if err != nil {
			return nil, err
		}
		desc.Tables[i].Columns = columns
	}

	return desc, nil
}
