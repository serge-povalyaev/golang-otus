package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunCmd(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		assert.Equal(t, successCode, RunCmd([]string{"ls"}, Environment{}))
	})

	t.Run("Error", func(t *testing.T) {
		assert.Equal(t, errorCode, RunCmd([]string{"cd tmp"}, Environment{}))
	})
}
