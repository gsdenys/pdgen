package options

import (
	"fmt"
	"sort"
	"strings"
)

// OutputOptions type that help work wit types
type OutputOptions string

// Options for output format
var Options = map[string]OutputOptions{
	"DEFAULT": OutputOptions(""),
	"MD":      OutputOptions("md"),
	"TXT":     OutputOptions("txt"),
	"CSV":     OutputOptions("csv"),
	"DOCX":    OutputOptions("docx"),
	"XLSX":    OutputOptions("xlsx"),
	"JSON":    OutputOptions("json"),
}

// getKeys returns an array of string containing the output possibilites
func getKeys() []string {
	keys := []string{}

	for k := range Options {
		keys = append(keys, strings.ToLower(k))
	}

	sort.Strings(keys)
	return keys
}

// String converts the object to string format
func (o *OutputOptions) String() string {
	return string(*o)
}

// Set provide a way to change de object value
func (o *OutputOptions) Set(v string) error {
	for k := range Options {
		if strings.ToLower(k) == v {
			*o = OutputOptions(v)
			return nil
		}
	}

	return fmt.Errorf(
		"must be one of %s",
		strings.Join(getKeys(), ", "),
	)
}

// Type returns the type of object
func (o *OutputOptions) Type() string {
	return "OutputOptions"
}

// Message returns the message that says the output possibilities
func Message() string {
	return fmt.Sprintf(
		`the possibles output. allowed: %s`,
		strings.Join(getKeys(), ", "),
	)
}
