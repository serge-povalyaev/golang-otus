package hw09structvalidator

import (
	"errors"
	"strconv"
	"strings"
)

var (
	minError   = errors.New("Значение должно быть больше")
	maxError   = errors.New("Значение должно быть меньше")
	inIntError = errors.New("Отсутствует в перечисленных вариантах")
)

func minValidator(value, min int) error {
	if value < min {
		return minError
	}

	return nil
}

func maxValidator(value, max int) error {
	if value > max {
		return maxError
	}

	return nil
}

func inIntValidator(value int, allowedValues string) error {
	if !contains(strings.Split(allowedValues, ","), strconv.Itoa(value)) {
		return inIntError
	}

	return nil
}
