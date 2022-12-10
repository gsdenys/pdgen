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
package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gsdenys/pdgen/pkg/models"
	_ "github.com/lib/pq"
)

func undefinedQueryError() error {
	return errors.New("undefined query error, try again or contact your DBA")
}

// This function will make a connection to the database only once.
func Connect(driver string, uri string) (*sql.DB, error) {
	var err error

	db, err := sql.Open(driver, uri)

	if err != nil {
		return nil, errors.New(`error when try to connect to the database"`)
	}

	if err = db.Ping(); err != nil {
		return nil, errors.New("connection was created but ping fail, so no content is accessible")
	}

	return db, nil
}

func GetDatabaseComment(db *sql.DB, database string) (string, error) {
	var desc string
	row := db.QueryRow(selectDatabaseComment, database)

	switch err := row.Scan(&desc); err {
	case nil:
		return desc, nil
	default:
		return "", fmt.Errorf("there no database named %s", database)
	}
}

func GetSchemaComment(db *sql.DB, schema string) (string, error) {
	var desc string
	row := db.QueryRow(selectSchemaComment, schema)
	switch err := row.Scan(&desc); err {
	case nil:
		return desc, nil
	default:
		return "", fmt.Errorf("there no schema named %s", schema)
	}
}

func GetAllTables(db *sql.DB, schema string) ([]models.Table, error) {
	var tbl []models.Table

	rows, err := db.Query(selectAllTables, schema)
	if err != nil {
		return nil, undefinedQueryError()
	}

	for rows.Next() {
		var table models.Table

		if err := rows.Scan(&table.Name, &table.Desc); err != nil {
			return nil, undefinedQueryError()
		}

		tbl = append(tbl, table)
	}

	return tbl, nil
}

func GetTableColumns(db *sql.DB, schema string, table string) ([]models.Columns, error) {
	var tbl []models.Columns

	rows, err := db.Query(selectTable, schema, table)
	if err != nil {
		return nil, undefinedQueryError()
	}

	for rows.Next() {
		var c models.Columns

		err := rows.Scan(&c.Column, &c.Type, &c.Allow, &c.Comment)
		if err != nil {
			return nil, undefinedQueryError()
		}

		tbl = append(tbl, c)
	}

	return tbl, nil
}
