package hw09structvalidator

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLenValidator(t *testing.T) {
	tests := []struct {
		value string
		len   int
		error error
	}{
		{
			value: "GOLANG",
			len:   6,
			error: nil,
		},
		{
			value: "GO",
			len:   6,
			error: errLen,
		},
		{
			value: "Го",
			len:   2,
			error: nil,
		},
	}

	for i, tc := range tests {
		tc := tc
		t.Run("Тест "+strconv.Itoa(i), func(t *testing.T) {
			require.True(t, errors.Is(lenValidator(tc.value, tc.len), tc.error))
		})
	}
}

func TestRegexpValidator(t *testing.T) {
	tests := []struct {
		value              string
		regexp             string
		validationError    error
		hasProcessingError bool
	}{
		{
			value:              "otus@otus.ru",
			regexp:             "^\\w+@\\w+\\.\\w+$",
			validationError:    nil,
			hasProcessingError: false,
		},
		{
			value:              "otus.ru",
			regexp:             "^\\w+@\\w+\\.\\w+$",
			validationError:    errRegexp,
			hasProcessingError: false,
		},
		{
			value:              "otus.ru",
			regexp:             "1\\2\\3...45",
			validationError:    nil,
			hasProcessingError: true,
		},
	}

	for i, tc := range tests {
		tc := tc
		t.Run("Тест "+strconv.Itoa(i), func(t *testing.T) {
			validationError, err := regexpValidator(tc.value, tc.regexp)
			require.True(t, errors.Is(validationError, tc.validationError))
			require.Equal(t, tc.hasProcessingError, err != nil)
		})
	}
}

func TestInStringValidator(t *testing.T) {
	tests := []struct {
		value string
		in    string
		error error
	}{
		{
			value: "GOLANG",
			in:    "GO,GOLANG",
			error: nil,
		},
		{
			value: "GO",
			in:    "GOLANG,PHP",
			error: errInString,
		},
		{
			value: "GO",
			in:    "",
			error: errInString,
		},
	}

	for i, tc := range tests {
		tc := tc
		t.Run("Тест "+strconv.Itoa(i), func(t *testing.T) {
			require.True(t, errors.Is(inStringValidator(tc.value, tc.in), tc.error))
		})
	}
}
