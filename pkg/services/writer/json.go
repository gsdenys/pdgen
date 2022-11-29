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

func (p *PrinterJson) Init(desc models.Describe) {
	// Do nothing because have nothing to initialise
}

func (p *PrinterJson) Title(title string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *PrinterJson) Subtitle(subtitle string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *PrinterJson) SubSubtitle(subSubtitle string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *PrinterJson) LineBreak() {
	//Do nothing because the unique action of this writer is Done
}

func (p *PrinterJson) Body(desc string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *PrinterJson) Columns(columns []models.Columns) {
	//Do nothing because the unique action of this writer is Done
}

func (p *PrinterJson) Table(t models.Table) {
	//Do nothing because the unique action of this writer is Done
}

func (p *PrinterJson) Done(desc models.Describe) {
	b, _ := json.MarshalIndent(desc, "", "    ")

	fmt.Fprintf(p.Out, "%s", string(b))

	_ = p.Out.(*os.File).Close()
}

func (p *PrinterJson) GetLanguage() *message.Printer {
	return p.Translate
}
