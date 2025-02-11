package helloworld

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	b := &bytes.Buffer{}
	HelloWorld(b)

	got := b.String()
	want := "HelloWorld"

	assert.Equal(t, want, got)
}
