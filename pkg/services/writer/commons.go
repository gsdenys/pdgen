package writer

import (
	"fmt"
	"io"
	"os"
)

func createFile(path string) io.Writer {
	file, err := os.Create(path)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return file
}
