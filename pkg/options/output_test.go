package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputOptions_Message(t *testing.T) {
	want := "the output types [default html json md txt]"

	assert.Equal(t, Message(), want)
}
