package lang

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestAmericanEnglish(t *testing.T) {
	lang := language.AmericanEnglish

	if got := AmericanEnglish(lang); !reflect.DeepEqual(got, lang) {
		t.Errorf("AmericanEnglish() = %v, want %v", got, lang)
	}

	printer := message.NewPrinter(lang)

	assert.Equal(t, printer.Sprintf("title"), "Data Dictionary")
}
