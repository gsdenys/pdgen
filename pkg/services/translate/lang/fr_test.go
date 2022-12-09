package lang

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestCanadianFrench(t *testing.T) {
	lang := language.CanadianFrench

	if got := CanadianFrench(lang); !reflect.DeepEqual(got, lang) {
		t.Errorf("CanadianFrench() = %v, want %v", got, lang)
	}

	printer := message.NewPrinter(lang)

	assert.Equal(t, printer.Sprintf("title"), "Dictionnaire de données")
}
