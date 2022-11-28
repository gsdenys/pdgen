package writer

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gsdenys/pdgen/pkg/models"
	"golang.org/x/text/message"
)

type PrinterJson struct {
	Out       io.Writer
	Translate *message.Printer
}

func (p *PrinterJson) Init(desc models.Describe) {}

func (p *PrinterJson) Title(title string) {}

func (p *PrinterJson) Subtitle(subtitle string) {}

func (p *PrinterJson) SubSubtitle(subSubtitle string) {}

func (p *PrinterJson) LineBreak() {}

func (p *PrinterJson) Body(desc string) {}

func (p *PrinterJson) Columns(columns []models.Columns) {}

func (p *PrinterJson) Table(t models.Table) {}

func (p *PrinterJson) Done(desc models.Describe) {
	b, err := json.MarshalIndent(desc, "", "    ")

	if err != nil {
		println("Error when try to convert to JSON!")
	}

	fmt.Fprintf(p.Out, "%s", string(b))

	_ = p.Out.(*os.File).Close()
}

func (p *PrinterJson) GetLanguage() *message.Printer {
	return p.Translate
}
