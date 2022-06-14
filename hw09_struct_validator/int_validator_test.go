package hw09structvalidator

import (
	"errors"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestMaxValidator(t *testing.T) {
	tests := []struct {
		value int
		max   int
		error error
	}{
		{
			value: 5,
			max:   6,
			error: nil,
		},
		{
			value: 5,
			max:   5,
			error: nil,
		},
		{
			value: 6,
			max:   5,
			error: maxError,
		},
	}

	for i, tc := range tests {
		tc := tc
		t.Run("Тест "+strconv.Itoa(i), func(t *testing.T) {
			require.True(t, errors.Is(maxValidator(tc.value, tc.max), tc.error))
		})
	}
}

func TestMinValidator(t *testing.T) {
	tests := []struct {
		value int
		min   int
		error error
	}{
		{
			value: 6,
			min:   5,
			error: nil,
		},
		{
			value: 5,
			min:   5,
			error: nil,
		},
		{
			value: 5,
			min:   6,
			error: minError,
		},
	}

	for i, tc := range tests {
		tc := tc
		t.Run("Тест "+strconv.Itoa(i), func(t *testing.T) {
			require.True(t, errors.Is(minValidator(tc.value, tc.min), tc.error))
		})
	}
}

func TestInIntValidator(t *testing.T) {
	tests := []struct {
		value int
		in    string
		error error
	}{
		{
			value: 1,
			in:    "1,2,3",
			error: nil,
		},
		{
			value: 1,
			in:    "1",
			error: nil,
		},
		{
			value: 1,
			in:    "2,3,4",
			error: inIntError,
		},
	}

	for i, tc := range tests {
		tc := tc
		t.Run("Тест "+strconv.Itoa(i), func(t *testing.T) {
			require.True(t, errors.Is(inIntValidator(tc.value, tc.in), tc.error))
		})
	}
}
