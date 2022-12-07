package writer

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/stretchr/testify/assert"
)

const wantColumns string = "\t<table class=\"styled-table\">\n\t\t<tr class=\"active-row\">\n\t\t\t<th>table-title-name</th>\n\t\t\t<th>table-title-type</th>\n\t\t\t<th>table-title-allow</th>\n\t\t\t<th>table-title-comment</th>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>name</td>\n\t\t\t<td>text</td>\n\t\t\t<td></td>\n\t\t\t<td>Somme comment</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>type</td>\n\t\t\t<td>text</td>\n\t\t\t<td>BASE EXTENSION</td>\n\t\t\t<td>Another comment</td>\n\t\t</tr>\n\t</table>\n"

var columns []models.Columns = []models.Columns{
	{
		Column:  "name",
		Type:    "text",
		Allow:   "",
		Comment: "Somme comment",
	},
	{
		Column:  "type",
		Type:    "text",
		Allow:   "BASE EXTENSION",
		Comment: "Another comment",
	},
}

func TestPrinterHTML_Init(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &HTML{
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

	p := &HTML{
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

	p := &HTML{
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

	p := &HTML{
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

	p := &HTML{
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

	p := &HTML{
		Out: createFile(file),
	}

	p.Body("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\t<p>test</p>\n")
}

func TestHTML_Columns(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	html := &HTML{
		Out: createFile(file),
	}

	translate.InitLanguage()
	html.Columns(columns)

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), wantColumns)
}

func TestHTML_SetWriter(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &HTML{}
	assert.Nil(t, p.Out)

	p.SetWriter(file)
	assert.NotNil(t, p.Out)
}

func TestHTML_Table(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	html := &HTML{
		Out: createFile(file),
	}

	translate.InitLanguage()

	table := models.Table{
		Name:    "test",
		Desc:    "Table test",
		Columns: columns,
	}

	html.Table(table)

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	want := "\t<h3>TEST</h3>\n\t<p>Table test</p>\n\t<br>\n" + wantColumns + "\t<br>\n"

	assert.Equal(t, string(f), want)
}

func TestHTML_Done(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	html := &HTML{
		Out: createFile(file),
	}

	translate.InitLanguage()
	html.Done(models.Describe{})

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\n</body>\n\n</html>")
}
