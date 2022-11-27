package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputOptions_Type(t *testing.T) {
	test := OutputOptions("md")

	assert.Equal(t, test.Type(), "OutputOptions")
}

func TestOutputOptions_String(t *testing.T) {
	test := OutputOptions("md")

	assert.Equal(t, test.String(), "md")
}

func TestOutputOptions_Message(t *testing.T) {
	want := "the possibles output. allowed: default, html, json, md, txt"

	assert.Equal(t, Message(), want)
}

func TestOutputOptions_Set(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		o       OutputOptions
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			o:       OutputOptions("md"),
			args:    args{v: "md"},
			wantErr: false,
		},
		{
			name:    "error",
			o:       OutputOptions("md"),
			args:    args{v: "nop"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.Set(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("OutputOptions.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
