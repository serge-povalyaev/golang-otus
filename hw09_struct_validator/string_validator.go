package hw09structvalidator

import (
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	lenError      = errors.New("Несоответствие длины строки")
	regexpError   = errors.New("Несоответствие регулярному выражению")
	inStringError = errors.New("Отсутствует в перечисленных вариантах")
)

func lenValidator(value string, length int) error {
	if utf8.RuneCountInString(value) != length {
		return lenError
	}

	return nil
}

func regexpValidator(value, regexpTemplate string) (error, error) {
	matched, err := regexp.MatchString(regexpTemplate, value)
	if err != nil {
		return nil, err
	}

	if !matched {
		return regexpError, nil
	}

	return nil, nil
}

func inStringValidator(value, allowedValues string) error {
	if !contains(strings.Split(allowedValues, ","), value) {
		return inStringError
	}

	return nil
}
