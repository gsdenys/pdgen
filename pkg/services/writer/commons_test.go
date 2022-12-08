package writer

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_createFile(t *testing.T) {
	var ok bool = false // The default value can be omitted :)

	exit = func(c int) {
		ok = true
	}

	createFile("/usr/bin/test.txt")

	assert.True(t, ok)
}

func Test_createFile_ok(t *testing.T) {
	file := getWorkDir() + uuid.NewString()

	createFile(file)

	_, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
}
