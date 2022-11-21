package database

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

const successConnection string = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

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
			if got := Connect(tt.args.driver, tt.args.uri); !reflect.DeepEqual(got == nil, tt.wantErr) {
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
				db:       Connect("postgres", successConnection),
				database: "postgres",
			},
			want: "standard public database",
		},
		{
			name: "comment-retrieved",
			args: args{
				db:       Connect("postgres", successConnection),
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
				db:     Connect("postgres", successConnection),
				schema: "public",
			},
			want: "standard public schema",
		},
		{
			name: "no-schema",
			args: args{
				db:     Connect("postgres", successConnection),
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
