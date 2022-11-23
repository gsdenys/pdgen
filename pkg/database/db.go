package database

import (
	"database/sql"
	"fmt"

	"github.com/gsdenys/pdgen/pkg/models"
	_ "github.com/lib/pq"
)

// This function will make a connection to the database only once.
func Connect(driver string, uri string) *sql.DB {
	var err error

	db, err := sql.Open(driver, uri)

	if err != nil {
		fmt.Printf("Connection error: %s", err.Error())
		return nil
	}

	if err = db.Ping(); err != nil {
		fmt.Printf("Connection established but ping fail: %s", err.Error())
		return nil
	}

	return db
}

func GetDatabaseComment(db *sql.DB, database string) (string, error) {
	var desc string
	row := db.QueryRow(selectDatabaseComment, database)
	switch err := row.Scan(&desc); err {
	case sql.ErrNoRows:
		return "", fmt.Errorf("there no database named %s", database)
	case nil:
		return desc, nil
	default:
		panic(err)
	}
}

func GetSchemaComment(db *sql.DB, schema string) (string, error) {
	var desc string
	row := db.QueryRow(selectSchemaComment, schema)
	switch err := row.Scan(&desc); err {
	case sql.ErrNoRows:
		return "", fmt.Errorf("there no schema named %s", schema)
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
		return nil, err
	}

	for rows.Next() {
		var table models.Table

		if err := rows.Scan(&table.Name, &table.Desc); err != nil {
			return nil, err
		}

		tbl = append(tbl, table)
	}

	return tbl, nil
}

func GetTableColumns(db *sql.DB, schema string, table string) ([]models.Columns, error) {
	var tbl []models.Columns

	rows, err := db.Query(selectTable, schema, table)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c models.Columns

		if err := rows.Scan(&c.Column, &c.Type, &c.Allow, &c.Comment); err != nil {
			return nil, err
		}

		tbl = append(tbl, c)
	}

	return tbl, nil
}
