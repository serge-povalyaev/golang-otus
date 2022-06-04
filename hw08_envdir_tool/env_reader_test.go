package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadDir(t *testing.T) {
	t.Run("dir not found", func(t *testing.T) {
		_, err := ReadDir("testdata/env/BAR")
		require.Error(t, ErrNotDir, err)
	})

	t.Run("env exists", func(t *testing.T) {
		env, _ := ReadDir("testdata/env")
		require.Equal(t, EnvValue{"bar", false}, env["BAR"])
		require.Equal(t, EnvValue{"", false}, env["EMPTY"])
		require.Equal(t, EnvValue{"   foo\nwith new line", false}, env["FOO"])
		require.Equal(t, EnvValue{"\"hello\"", false}, env["HELLO"])
		_, ok := env["UNSET"]
		require.Equal(t, false, ok)
		_, ok = env["EQUAL=EQUAL"]
		require.Equal(t, false, ok)
		require.Equal(t, EnvValue{"123", false}, env["TAB"])
	})
}
