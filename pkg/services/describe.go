package services

import (
	"github.com/gsdenys/pdgen/pkg/database"
	"github.com/gsdenys/pdgen/pkg/models"
)

func Describe(uri string, db string) (*models.Describe, error) {
	desc := &models.Describe{}

	conn := database.Connect("postgres", uri)
	defer conn.Close()

	dbDesc, err := database.GetDatabaseComment(conn, db)
	if err != nil {
		return nil, err
	}

	desc.Database = dbDesc

	return desc, nil
}
