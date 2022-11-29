package writer

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/stretchr/testify/assert"
)

func TestPrinterHTML_Init(t *testing.T) {
	file := os.TempDir() + uuid.NewString()

	p := &PrinterHTML{
		Out:       createFile(file),
		Translate: translate.GetTranslation("en"),
	}

	p.Init(models.Describe{})

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), Base+"\n")
}

func TestPrinterHTML_Title(t *testing.T) {
	file := os.TempDir() + uuid.NewString()

	p := &PrinterHTML{
		Out:       createFile(file),
		Translate: translate.GetTranslation("en"),
	}

	p.Title("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<h1>TEST</h1>\n")
}

func TestPrinterHTML_Subtitle(t *testing.T) {
	file := os.TempDir() + uuid.NewString()

	p := &PrinterHTML{
		Out:       createFile(file),
		Translate: translate.GetTranslation("en"),
	}

	p.Subtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<h2>TEST</h2>\n")
}

func TestPrinterHTML_SubSubtitle(t *testing.T) {
	file := os.TempDir() + uuid.NewString()

	p := &PrinterHTML{
		Out:       createFile(file),
		Translate: translate.GetTranslation("en"),
	}

	p.SubSubtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<h3>TEST</h3>\n")
}
