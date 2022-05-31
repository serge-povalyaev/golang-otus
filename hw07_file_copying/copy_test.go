package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	t.Run("file not found", func(t *testing.T) {
		err := Copy("testdata/404.txt", "", 0, 0, 1)
		require.Error(t, ErrFileNotFound, err)
	})

	t.Run("unsupported", func(t *testing.T) {
		err := Copy("testdata", "", 0, 0, 1)
		require.Equal(t, ErrUnsupportedFile, err)

		err = Copy("testdata/0.txt", "", 0, 0, 1)
		require.Error(t, ErrUnsupportedFile, err)

		err = Copy("testdata/", "", 0, 0, 1)
		require.Error(t, ErrUnsupportedFile, err)

		err = Copy("testdata/input.txt", "testdata/", 0, 0, 1)
		require.Error(t, os.ErrInvalid, err)

		err = Copy("testdata/input.txt", "testdata1/out.txt", 0, 0, 1)
		require.Error(t, os.ErrInvalid, err)
	})

	t.Run("offset error", func(t *testing.T) {
		err := Copy("testdata/input.txt", "", 10000, 0, 1)
		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})

	t.Run("bar size", func(t *testing.T) {
		size := getBarSize(1000, 900, 90)
		require.Equal(t, 90, size)

		size = getBarSize(1000, 950, 90)
		require.Equal(t, 50, size)
	})

	t.Run("buf size zero", func(t *testing.T) {
		err := Copy("testdata/tmp.txt", "", 0, 0, 0)
		require.Error(t, ErrBufSize, err)
	})
}
