package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const selectDatabaseComment string = `
SELECT pg_catalog.shobj_description(d.oid, 'pg_database') AS "description"
FROM pg_catalog.pg_database d 
WHERE  datname = $1`

const selectSchemaComment string = `SELECT obj_description($1::regnamespace, 'pg_namespace')`

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
