package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ttt()
	assert.Equal(t, 1, 1, "they should be equal")
}
