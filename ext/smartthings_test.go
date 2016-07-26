package ext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStackMalformedInput(t *testing.T) {
	stack, err := ParseStack("foobar")

	assert.Error(t, err)
	assert.Nil(t, stack)
}

func TestParseStackUnmappedEnvironment(t *testing.T) {
	stack, err := ParseStack("na01f_useast1")

	assert.Error(t, err)
	assert.Nil(t, stack)
}

func TestParseStackValid(t *testing.T) {
	stack, err := ParseStack("na01d_uswest2")

	assert.Nil(t, err)
	assert.Equal(t, "na01d", stack.Name)
	assert.Equal(t, "dev", stack.Environment)
	assert.Equal(t, "uswest2", stack.Region)
}
