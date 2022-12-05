package database

import (
	"database/sql"
	"errors"

	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	_ "github.com/lib/pq"
)

// This function will make a connection to the database only once.
func Connect(driver string, uri string) (*sql.DB, error) {
	var err error

	db, err := sql.Open(driver, uri)

	if err != nil {
		return nil, errors.New(translate.T.Sprintf("connect-error", uri))
	}

	if err = db.Ping(); err != nil {
		return nil, errors.New(translate.T.Sprintf("ping-error"))
	}

	return db, nil
}

func GetDatabaseComment(db *sql.DB, database string) (string, error) {
	var desc string
	row := db.QueryRow(selectDatabaseComment, database)

	switch err := row.Scan(&desc); err {
	case sql.ErrNoRows:
		return "", errors.New(translate.T.Sprintf("db-not-found", database))
	case nil:
		return desc, nil
	default:
		return "", errors.New(translate.T.Sprintf("undefined-error"))
	}
}

func GetSchemaComment(db *sql.DB, schema string) (string, error) {
	var desc string
	row := db.QueryRow(selectSchemaComment, schema)
	switch err := row.Scan(&desc); err {
	case sql.ErrNoRows:
		return "", errors.New(translate.T.Sprintf("schema-not-found", schema))
	case nil:
		return desc, nil
	default:
		return "", errors.New(translate.T.Sprintf("undefined-error"))
	}
}

func GetAllTables(db *sql.DB, schema string) ([]models.Table, error) {
	var tbl []models.Table

	rows, err := db.Query(selectAllTables, schema)
	if err != nil {
		return nil, errors.New(translate.T.Sprintf("db-undefined-error"))
	}

	for rows.Next() {
		var table models.Table

		if err := rows.Scan(&table.Name, &table.Desc); err != nil {
			return nil, errors.New(translate.T.Sprintf("db-undefined-error"))
		}

		tbl = append(tbl, table)
	}

	return tbl, nil
}

func GetTableColumns(db *sql.DB, schema string, table string) ([]models.Columns, error) {
	var tbl []models.Columns

	rows, err := db.Query(selectTable, schema, table)
	if err != nil {
		return nil, errors.New(translate.T.Sprintf("db-undefined-error"))
	}

	for rows.Next() {
		var c models.Columns

		if err := rows.Scan(&c.Column, &c.Type, &c.Allow, &c.Comment); err != nil {
			return nil, errors.New(translate.T.Sprintf("undefined-error"))
		}

		tbl = append(tbl, c)
	}

	return tbl, nil
}
