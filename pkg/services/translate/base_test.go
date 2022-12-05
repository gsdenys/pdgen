package translate

import (
	"reflect"
	"testing"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestGetTranslation(t *testing.T) {
	type args struct {
		l string
	}
	tests := []struct {
		name string
		args args
		want *message.Printer
	}{
		{
			name: "português",
			args: args{
				l: "pt",
			},
			want: message.NewPrinter(language.BrazilianPortuguese),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetTranslation(tt.args.l)
			if !reflect.DeepEqual(got, tt.want) {
				got.Print("Data Dictionary for database")
				t.Errorf("GetTranslation() = %v, want %v", got, tt.want)
			}
			got.Print("Data Dictionary for database")
		})
	}
}
