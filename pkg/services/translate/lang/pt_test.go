package lang

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestBrazilianPortuguese(t *testing.T) {
	lang := language.BrazilianPortuguese

	if got := BrazilianPortuguese(lang); !reflect.DeepEqual(got, lang) {
		t.Errorf("BrazilianPortuguese() = %v, want %v", got, lang)
	}

	printer := message.NewPrinter(lang)

	assert.Equal(t, printer.Sprintf("title"), "Dicionário de Dados")
}
