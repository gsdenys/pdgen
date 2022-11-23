package models

import (
	"testing"
)

func TestBasic_String(t *testing.T) {
	type fields struct {
		Name string
		Desc string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "full-data",
			fields: fields{
				Name: "test",
				Desc: "This is some test",
			},
			want: "TEST\nThis is some test\n\n",
		},
		{
			name: "no-desc",
			fields: fields{
				Name: "test",
				Desc: "",
			},
			want: "TEST\n\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Basic{
				Name: tt.fields.Name,
				Desc: tt.fields.Desc,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("Basic.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_String(t *testing.T) {
	type fields struct {
		Name    string
		Desc    string
		Columns []Columns
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "full content",
			fields: fields{
				Name: "test",
				Desc: "This is some test",
				Columns: []Columns{
					{
						Column:  "id",
						Type:    "integer",
						Allow:   "",
						Comment: "Unique identifier for this table",
					},
					{
						Column:  "name",
						Type:    "text",
						Allow:   "",
						Comment: "the name of test",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				Name:    tt.fields.Name,
				Desc:    tt.fields.Desc,
				Columns: tt.fields.Columns,
			}
			if got := tr.String(); got != tt.want {
				t.Errorf("Table.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
