package writer

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/stretchr/testify/assert"
)

const wantColumnsMd = `| table-title-name | table-title-type | table-title-allow | table-title-comment |
| :--- | :--- | :----: | :--- |
| name | text |  | Somme comment |
| type | text | BASE EXTENSION | Another comment |
`

func TestMD_SetWriter(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &MD{}
	assert.Nil(t, p.Out)

	p.SetWriter(file)
	assert.NotNil(t, p.Out)
}

func TestMD_Init(t *testing.T) {
	p := &MD{}
	assert.Nil(t, p.Out)

	p.Init(models.Describe{})
	assert.Nil(t, p.Out)
}

func TestMD_Title(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &MD{
		Out: CreateFile(file),
	}

	p.Title("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "# TEST\n")
}

func TestMD_Subtitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &MD{
		Out: CreateFile(file),
	}

	p.Subtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "## TEST\n")
}

func TestMD_SubSubtitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &MD{
		Out: CreateFile(file),
	}

	p.SubSubtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "### TEST\n")
}

func TestMD_Body(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &MD{
		Out: CreateFile(file),
	}

	p.Body("Some test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "Some test\n")
}

func TestMD_LineBreak(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &MD{
		Out: CreateFile(file),
	}

	p.LineBreak()

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\n")
}

func TestMD_Columns(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	html := &MD{
		Out: CreateFile(file),
	}

	translate.InitLanguage()
	html.Columns(columns)

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), wantColumnsMd)
}

func TestMD_Table(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	md := &MD{
		Out: CreateFile(file),
	}

	translate.InitLanguage()

	table := models.Table{
		Name:    "test",
		Desc:    "Table test",
		Columns: columns,
	}

	md.Table(table)

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	want := "### TEST\nTable test\n\n" + wantColumnsMd + "\n"

	assert.Equal(t, string(f), want)
}

func TestMD_Done(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &MD{
		Out: CreateFile(file),
	}

	p.Done(models.Describe{})
}
