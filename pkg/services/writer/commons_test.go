package writer

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_createFile(t *testing.T) {
	file, err := CreateFile("/usr/bin/test.txt")

	assert.Nil(t, file)
	assert.Error(t, err)
}

func Test_createFile_ok(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	_, _ = CreateFile(file)

	_, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
}
