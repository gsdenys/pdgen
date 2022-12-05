package writer

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestPrinterHTML_Init(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &PrinterHTML{
		Out: createFile(file),
	}

	p.Init(models.Describe{})

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), Base+"\n")
}

func TestPrinterHTML_Title(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &PrinterHTML{
		Out: createFile(file),
	}

	p.Title("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<h1>TEST</h1>\n")
}

func TestPrinterHTML_Subtitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &PrinterHTML{
		Out: createFile(file),
	}

	p.Subtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<h2>TEST</h2>\n")
}

func TestPrinterHTML_SubSubtitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &PrinterHTML{
		Out: createFile(file),
	}

	p.SubSubtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<h3>TEST</h3>\n")
}

func TestPrinterHTML_LineBreak(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &PrinterHTML{
		Out: createFile(file),
	}

	p.LineBreak()

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<br>\n")
}

func TestPrinterHTML_Body(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &PrinterHTML{
		Out: createFile(file),
	}

	p.Body("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<p>test</p>\n")
}

// func TestPrinterHTML_Column(t *testing.T) {
// 	file := getWorkDir() + uuid.NewString()

// 	p := &PrinterHTML{
// 		Out:       createFile(file),
// 		Translate: translate.GetTranslation("en"),
// 	}

// 	p.Body("test")

// 	f, err := os.ReadFile(file)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	assert.Equal(t, string(f), "\t<p>TEST</p>\n")
// }
