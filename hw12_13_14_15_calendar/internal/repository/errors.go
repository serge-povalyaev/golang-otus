package repository

import (
	"errors"
)

var (
	ErrEventNotFound = errors.New("событие не найдено")
	ErrEventExists   = errors.New("событие с указанным идентификатором существует")
)
