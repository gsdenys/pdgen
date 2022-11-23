package services

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gsdenys/pdgen/pkg/models"
)

func ToJSON(desc models.Describe, path string) bool {
	b, err := json.MarshalIndent(desc, "", "    ")

	if err != nil {
		println("Error when try to convert to JSON!")
		return false
	}

	err = ioutil.WriteFile(path, b, 0644)

	if err != nil {
		println("Error when try to save JSON file!")
		println(string(b))
		return false
	}

	return true
}
