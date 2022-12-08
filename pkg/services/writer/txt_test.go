package writer

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/stretchr/testify/assert"
)

func TestTXT_SetWriter(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &TXT{}
	assert.Nil(t, p.Out)

	p.SetWriter(file)
	assert.NotNil(t, p.Out)
}

func TestTXT_Init(t *testing.T) {
	p := &TXT{}
	assert.Nil(t, p.Out)

	p.Init(models.Describe{})
	assert.Nil(t, p.Out)
}

func TestTXT_Title(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &TXT{
		Out: createFile(file),
	}

	p.Title("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "TEST\n")
}

func TestTXT_Subtitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &TXT{
		Out: createFile(file),
	}

	p.Subtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "TEST\n")
}

func TestTXT_SubSubTitle(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &TXT{
		Out: createFile(file),
	}

	p.SubSubtitle("test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "TEST\n")
}

func TestTXT_LineBreak(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &TXT{
		Out: createFile(file),
	}

	p.LineBreak()

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "\n")
}

func TestTXT_Body(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &TXT{
		Out: createFile(file),
	}

	p.Body("Some test")

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), "Some test\n")
}

func TestTXT_Columns(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &TXT{
		Out: createFile(file),
	}

	//columns is defined at html_test.go
	translate.InitLanguage()
	p.Columns(columns)

	f, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(f), wantColumnsConsole)
}

func TestTXT_Table(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	md := &TXT{
		Out: createFile(file),
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

	want := "TEST\nTable test\n\n" + wantColumnsConsole + "\n"

	assert.Equal(t, string(f), want)
}

func TestTXT_Done(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	p := &TXT{
		Out: createFile(file),
	}

	p.Done(models.Describe{})
}
