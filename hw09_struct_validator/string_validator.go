package hw09structvalidator

import (
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	errLen      = errors.New("несоответствие длины строки")
	errRegexp   = errors.New("несоответствие регулярному выражению")
	errInString = errors.New("отсутствует в перечисленных вариантах")
)

func lenValidator(value string, length int) error {
	if utf8.RuneCountInString(value) != length {
		return errLen
	}

	return nil
}

func regexpValidator(value, regexpTemplate string) (error, error) {
	matched, err := regexp.MatchString(regexpTemplate, value)
	if err != nil {
		return nil, err
	}

	if !matched {
		return errRegexp, nil
	}

	return nil, nil
}

func inStringValidator(value, allowedValues string) error {
	if !contains(strings.Split(allowedValues, ","), value) {
		return errInString
	}

	return nil
}
