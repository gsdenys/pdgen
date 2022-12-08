package writer

import (
	"fmt"
	"io"
	"os"
)

var exit func(code int) = os.Exit

func CreateFile(path string) io.Writer {
	file, err := os.Create(path)

	if err != nil {
		fmt.Println(err.Error())
		exit(1)
	}

	return file
}
