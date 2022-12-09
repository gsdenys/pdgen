package writer

import (
	"io"
	"os"
)

func CreateFile(path string) (io.Writer, error) {
	file, err := os.Create(path)

	if err != nil {
		return nil, err
	}

	return file, nil
}
