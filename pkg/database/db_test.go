package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/gsdenys/pdgen/pkg/models"
	_ "github.com/lib/pq"
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

func connError(driver string, uri string) *sql.DB {
	conn, _ := Connect(driver, uri)
	return conn
}

func TestConnect(t *testing.T) {
	type args struct {
		driver string
		uri    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "connection-successful",
			args: args{
				driver: "postgres",
				uri:    "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
			},
			wantErr: false,
		},
		{
			name: "connection-error",
			args: args{
				driver: "bla",
				uri:    "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
			},
			wantErr: true,
		},
		{
			name: "ping-error",
			args: args{
				driver: "postgres",
				uri:    "postgres://idontknow:nitherami@localhost:4321/some",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connError(tt.args.driver, tt.args.uri); !reflect.DeepEqual(got == nil, tt.wantErr) {
				t.Errorf("Connect() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestGetDatabaseComment(t *testing.T) {
	type args struct {
		db       *sql.DB
		database string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "comment-retrieved",
			args: args{
				db:       connError("postgres", successConnection),
				database: "postgres",
			},
			want: "standard public database",
		},
		{
			name: "comment-not-retrieved",
			args: args{
				db:       connError("postgres", successConnection),
				database: "bla",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetDatabaseComment(tt.args.db, tt.args.database); got != tt.want {
				t.Errorf("GetDatabaseComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSchemaComment(t *testing.T) {
	type args struct {
		db     *sql.DB
		schema string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "comment-retrieved",
			args: args{
				db:     connError("postgres", successConnection),
				schema: "public",
			},
			want: "standard public schema",
		},
		{
			name: "no-schema",
			args: args{
				db:     connError("postgres", successConnection),
				schema: "bla",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSchemaComment(tt.args.db, tt.args.schema)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSchemaComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSchemaComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllTables(t *testing.T) {

	cnn := connError("postgres", successConnection)
	if _, err := cnn.Exec(script); err != nil {
		t.Errorf("create table and comments error: %s", err.Error())
	}
	cnn.Close()

	cnn = connError("postgres", successConnection)
	res, _ := GetAllTables(cnn, "public")

	if res == nil {
		t.Error("No content found")
	}

	for i := range res {
		if res[i].Name == "test" && res[i].Desc == "table for test propose" {
			return
		}
	}

	t.Error("table test not found")
}

func TestGetTableColumns(t *testing.T) {
	cnn := connError("postgres", successConnection)
	if _, err := cnn.Exec(script); err != nil {
		t.Errorf("create table and comments error: %s", err.Error())
	}
	cnn.Close()

	type args struct {
		db     *sql.DB
		schema string
		table  string
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Columns
		wantErr bool
	}{
		{
			name: "success-retrieve-data",
			args: args{
				db:     connError("postgres", successConnection),
				schema: "public",
				table:  "test",
			},
			want: []models.Columns{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTableColumns(tt.args.db, tt.args.schema, tt.args.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTableColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTableColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}
