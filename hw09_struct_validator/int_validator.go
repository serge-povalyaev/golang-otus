package hw09structvalidator

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errMin   = errors.New("значение должно быть больше")
	errMax   = errors.New("значение должно быть меньше")
	errInInt = errors.New("отсутствует в перечисленных вариантах")
)

func minValidator(value, min int) error {
	if value < min {
		return errMin
	}

	return nil
}

func maxValidator(value, max int) error {
	if value > max {
		return errMax
	}

	return nil
}

func inIntValidator(value int, allowedValues string) error {
	if !contains(strings.Split(allowedValues, ","), strconv.Itoa(value)) {
		return errInInt
	}

	return nil
}
