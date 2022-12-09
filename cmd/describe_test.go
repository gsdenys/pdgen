/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/gsdenys/pdgen/pkg/services"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/gsdenys/pdgen/pkg/services/writer"
	"github.com/stretchr/testify/assert"
)

func Test_setLang(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "no-lang-selected",
			args:    args{lang: ""},
			wantErr: false,
		},
		{
			name:    "pt-BR",
			args:    args{lang: "pt-BR"},
			wantErr: false,
		},
		{
			name:    "te-TE",
			args:    args{lang: "te-TE"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		translate.Register()

		var ok bool = false

		exit = func(code int) {
			ok = true
		}

		t.Run(tt.name, func(t *testing.T) {
			setLang(tt.args.lang)

			if tt.wantErr {
				assert.True(t, ok)
			}
		})
	}
}

func Test_getFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name    string
		args    args
		want    services.Printer
		wantErr bool
	}{
		{
			name:    "txt",
			args:    args{format: "txt"},
			want:    &writer.TXT{},
			wantErr: false,
		},
		{
			name:    "bla",
			args:    args{format: "bla"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got, err := getFormat(tt.args.format); !reflect.DeepEqual(got, tt.want) && (err != nil) == tt.wantErr {
				t.Errorf("getFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_describeCmd(t *testing.T) {
	type args struct {
		actual *bytes.Buffer
		arg    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				actual: bytes.NewBuffer([]byte{}),
				arg:    []string{"describe"},
			},
			want: "",
		},
		{
			name: "formatNotSupported",
			args: args{
				actual: bytes.NewBuffer([]byte{}),
				arg:    []string{"describe", "-fexe"},
			},
			want: "the format exe is not acceptable, please select one of: [default html json md txt]",
		},
		{
			name: "permissionDenied",
			args: args{
				actual: bytes.NewBuffer([]byte{}),
				arg:    []string{"describe", "-ftxt", "-o/usr/bin/test.txt"},
			},
			want: "open /usr/bin/test.txt: permission denied",
		},
		{
			name: "dbNotFound",
			args: args{
				actual: new(bytes.Buffer),
				arg:    []string{"describe", "-dtest", "-o "},
			},
			want: "there no database named test",
		},
		{
			name: "schemaNotFound",
			args: args{
				actual: new(bytes.Buffer),
				arg:    []string{"describe", "-sasdfg", "-dpostgres", "-o "},
			},
			want: "there no schema named asdfg",
		},
		{
			name: "file-generated-txt",
			args: args{
				actual: new(bytes.Buffer),
				arg:    []string{"describe", "-spublic", "-dpostgres", "-o ", "-ftxt"},
			},
			want: "TXT document created.\n",
		},
		{
			name: "file-generated-html",
			args: args{
				actual: new(bytes.Buffer),
				arg:    []string{"describe", "-spublic", "-dpostgres", "-o ", "-fhtml"},
			},
			want: "HTML document created.\n",
		},
		{
			name: "file-generated-md",
			args: args{
				actual: new(bytes.Buffer),
				arg:    []string{"describe", "-spublic", "-dpostgres", "-o ", "-fmd"},
			},
			want: "MD document created.\n",
		},
	}
	for _, tt := range tests {
		translate.InitLanguage()
		t.Run(tt.name, func(t *testing.T) {
			rootCmd.SetOut(tt.args.actual)
			rootCmd.SetErr(tt.args.actual)

			rootCmd.SetArgs(tt.args.arg)

			err := rootCmd.Execute()
			if err != nil {
				t.Error(err)
			}

			fmt.Printf("%v", tt.args.arg)

			tt.args.arg = nil

			assert.Equal(t, tt.args.actual.String(), tt.want)
		})
	}
}
