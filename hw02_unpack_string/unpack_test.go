package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Стандартное поведение
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		// Тестирование без цифр
		{input: "abccd", expected: "abccd"},
		// Проверка пустой строки
		{input: "", expected: ""},
		// Проверка использования нуля
		{input: "aaa0b", expected: "aab"},
		// Проверка дополнительных символов
		{input: "\n3bc0", expected: "\n\n\nb"},
		// Проверка кириллицы
		{input: "ПП0р1и2в3е4т", expected: "Приивввеееет"},
		// Проверка использования пробелов
		{input: "O 2T 3U 0S", expected: "O  T   US"},
		// Проверка экранирования обратного слеша
		{input: "\\3", expected: "\\\\\\"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"5", "3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
