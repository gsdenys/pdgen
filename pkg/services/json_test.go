package services

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/gsdenys/pdgen/pkg/models"
)

func getTempDir() string {
	dir, err := ioutil.TempDir("test", "")
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func TestToJSON(t *testing.T) {
	type args struct {
		desc models.Describe
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "output-successful",
			args: args{
				desc: models.Describe{
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
				path: os.TempDir() + "/" + uuid.NewString() + ".json",
			},
			want: true,
		},

		{
			name: "permission-deny",
			args: args{
				desc: models.Describe{},
				path: "/root/" + uuid.NewString() + ".json",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToJSON(tt.args.desc, tt.args.path); got != tt.want {
				t.Errorf("ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
