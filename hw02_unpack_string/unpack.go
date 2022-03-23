package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var builder strings.Builder
	var prev rune
	for _, current := range str {
		// Запрещаем первому символу быть числом
		if prev == 0 && unicode.IsDigit(current) {
			return "", ErrInvalidString
		}

		// Делаем для того, чтобы не добавлять нулевой символ в строку
		if prev == 0 {
			prev = current
			continue
		}

		// Запрещаем числа друг за другом
		if unicode.IsDigit(current) && unicode.IsDigit(prev) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(current) {
			// Повторяем символ укаазанное количество раз
			count, e := strconv.Atoi(string(current))
			if e != nil {
				return "", ErrInvalidString
			}
			builder.WriteString(strings.Repeat(string(prev), count))
		} else if !unicode.IsDigit(prev) {
			// Если символы идут друг за другом, то повторяем предыдущий символ 1 раз
			// (т.к. число повторов не указано)
			builder.WriteString(string(prev))
		}

		prev = current
	}

	// На случай, если последний элемент стороки - символ
	if !unicode.IsDigit(prev) && prev != 0 {
		builder.WriteString(string(prev))
	}

	return builder.String(), nil
}
