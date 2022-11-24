package services

import (
	"bytes"
	"io/ioutil"

	"github.com/gsdenys/pdgen/pkg/models"
)

func ToTXT(desc models.Describe, path string, lang string) bool {
	var out bytes.Buffer
	printer := &PrinterText{
		out:      &out,
		language: lang,
	}

	printer.print(desc)

	err := ioutil.WriteFile(path, out.Bytes(), 0644)

	if err != nil {
		println("Error when try to save JSON file!")
		println(out.String())
		return false
	}

	return true

}
