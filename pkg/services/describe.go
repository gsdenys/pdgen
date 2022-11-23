package services

import (
	"github.com/gsdenys/pdgen/pkg/database"
	"github.com/gsdenys/pdgen/pkg/models"
)

func Describe(uri string, db string, schema string) (*models.Describe, error) {
	desc := &models.Describe{}

	conn := database.Connect("postgres", uri)
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

	//get tables
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
