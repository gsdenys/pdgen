package services

import (
	"reflect"
	"testing"

	"github.com/gsdenys/pdgen/pkg/database"
	"github.com/gsdenys/pdgen/pkg/models"
)

const successConnection string = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

const script string = `
	CREATE TABLE IF NOT EXISTS test (
		id serial PRIMARY KEY,
		name text NOT NULL
   );
   COMMENT ON TABLE test IS 'table for test propose';
   
   COMMENT ON COLUMN test.id IS 'sequencial unique identifier';
   COMMENT ON COLUMN test.name IS 'name of test';
	`

func TestDescribe(t *testing.T) {

	cnn := database.Connect("postgres", successConnection)
	if _, err := cnn.Exec(script); err != nil {
		t.Errorf("create table and comments error: %s", err.Error())
	}
	cnn.Close()

	type args struct {
		uri    string
		db     string
		schema string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Describe
		wantErr bool
	}{
		{
			name: "success-retrieved-data",
			args: args{
				uri:    successConnection,
				db:     "postgres",
				schema: "public",
			},
			want: &models.Describe{
				Database: models.Basic{
					Name: "postgres",
					Desc: "standard public database",
				},
				Schema: models.Basic{
					Name: "public",
					Desc: "standard public schema",
				},
				Tables: []models.Table{
					{
						Name: "test",
						Desc: "table for test propose",
						Columns: []models.Columns{
							{
								Column:  "id",
								Type:    "integer",
								Allow:   "",
								Comment: "sequencial unique identifier",
							},
							{
								Column:  "name",
								Type:    "text",
								Allow:   "",
								Comment: "name of test",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Describe(tt.args.uri, tt.args.db, tt.args.schema)
			if (err != nil) != tt.wantErr {
				t.Errorf("Describe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Describe() = %v, want %v", got, tt.want)
			}
		})
	}
}
