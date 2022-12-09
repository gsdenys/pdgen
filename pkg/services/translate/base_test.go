package translate

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestRegister(t *testing.T) {
	tests := []struct {
		name string
		lang language.Tag
		want string
	}{
		{
			name: "pt-BR",
			lang: language.BrazilianPortuguese,
			want: "Dicionário de Dados",
		},
		{
			name: "en-US",
			lang: language.AmericanEnglish,
			want: "Data Dictionary",
		},
		{
			name: "fr-CA",
			lang: language.CanadianFrench,
			want: "Dictionnaire de données",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register()

			printer := message.NewPrinter(tt.lang)
			assert.Equal(t, printer.Sprintf("title"), tt.want)
		})
	}
}

func TestSetLanguage(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "en-US",
			args: args{lang: language.AmericanEnglish.String()},
			want: true,
		},
		{
			name: "nt-EX",
			args: args{lang: "nt-EX"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register()
			if got := SetLanguage(tt.args.lang); got != tt.want {
				t.Errorf("SetLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitLanguage(t *testing.T) {
	Register()
	InitLanguage()
}

func TestGetKeys(t *testing.T) {

	var want []string = []string{
		"en",
		"en-US",
		"fr",
		"fr-CA",
		"pt",
		"pt-BR",
	}

	if got := GetKeys(); !reflect.DeepEqual(got, want) {
		t.Errorf("GetKeys() = %v, want %v", got, want)
	}
}
