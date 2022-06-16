package hw09structvalidator

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

var (
	ErrUnsupportedValidator = errors.New("неподдерживаемый валидатор")
	ErrUnsupportedType      = errors.New("неподдерживаемый тип данных")
)

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	result := make([]string, 0)
	for _, err := range v {
		result = append(result, err.Field+": "+err.Err.Error())
	}

	return strings.Join(result, "\n")
}

func Validate(v interface{}) error {
	var validationErrors ValidationErrors
	info := reflect.ValueOf(v)
	if info.Kind() != reflect.Struct {
		return ErrUnsupportedType
	}

	infoType := info.Type()

	for i := 0; i < info.NumField(); i++ {
		fieldInfo := infoType.Field(i)

		tagValue := fieldInfo.Tag.Get("validate")
		if tagValue == "" {
			continue
		}

		validatorsSlice := strings.Split(tagValue, "|")
		fieldValue := info.Field(i)
		for _, validator := range validatorsSlice {
			validatorInfo := strings.Split(validator, ":")
			validatorName := validatorInfo[0]
			validatorValue := validatorInfo[1]
			fieldTypeKind := fieldInfo.Type.Kind()

			if fieldTypeKind == reflect.Slice {
				sliceValidationErrors, err := validateSlice(ValidateData{validatorName, validatorValue, fieldValue}, fieldInfo.Name)
				if err != nil {
					return err
				}

				validationErrors = append(validationErrors, sliceValidationErrors...)
			} else {
				validationError, err := validateValue(ValidateData{validatorName, validatorValue, fieldValue}, fieldTypeKind)
				if err != nil {
					return err
				}

				if validationError != nil {
					validationErrors = append(validationErrors, ValidationError{fieldInfo.Name, validationError})
				}
			}
		}
	}

	if len(validationErrors) == 0 {
		return nil
	}

	return validationErrors
}

type ValidateData struct {
	validatorName  string
	validatorValue string
	fieldValue     reflect.Value
}

func validateValue(data ValidateData, fieldTypeKind reflect.Kind) (error, error) {
	switch fieldTypeKind { //nolint:exhaustive
	case reflect.String:
		return data.ValidateString()

	case reflect.Int:
		return data.ValidateInt()

	default:
		return nil, ErrUnsupportedType
	}
}

func validateSlice(data ValidateData, fieldName string) (ValidationErrors, error) {
	var validationErrors ValidationErrors
	slice := data.fieldValue.Slice(0, data.fieldValue.Len())
	sliceOf := data.fieldValue.Type().Elem().Kind()
	for i := 0; i < data.fieldValue.Len(); i++ {
		var validationError error
		var err error
		switch sliceOf { //nolint:exhaustive
		case reflect.String, reflect.Int:
			validationError, err = validateValue(ValidateData{data.validatorName, data.validatorValue, slice.Index(i)}, sliceOf)
			if err != nil {
				return nil, err
			}
		default:
			return nil, ErrUnsupportedType
		}

		if validationError != nil {
			validationErrors = append(
				validationErrors,
				ValidationError{fieldName + "[" + strconv.Itoa(i) + "]", validationError},
			)
		}
	}

	return validationErrors, nil
}

func (data ValidateData) ValidateString() (error, error) {
	switch data.validatorName {
	case "len":
		num, err := strconv.Atoi(data.validatorValue)
		if err != nil {
			return nil, err
		}

		return lenValidator(data.fieldValue.String(), num), nil

	case "regexp":
		return regexpValidator(data.fieldValue.String(), data.validatorValue)

	case "in":
		return inStringValidator(data.fieldValue.String(), data.validatorValue), nil

	default:
		return nil, ErrUnsupportedValidator
	}
}

func (data ValidateData) ValidateInt() (error, error) {
	switch data.validatorName {
	case "min":
		num, err := strconv.Atoi(data.validatorValue)
		if err != nil {
			return nil, err
		}

		return minValidator(int(data.fieldValue.Int()), num), nil

	case "max":
		num, err := strconv.Atoi(data.validatorValue)
		if err != nil {
			return nil, err
		}

		return maxValidator(int(data.fieldValue.Int()), num), nil

	case "in":
		return inIntValidator(int(data.fieldValue.Int()), data.validatorValue), nil

	default:
		return nil, ErrUnsupportedValidator
	}
}
