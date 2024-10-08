package utils

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func TrimSpacesInStruct(s interface{}) {
	v := reflect.ValueOf(s).Elem()

	// Tüm alanları dolaşarak sadece string türünde olanları kontrol ederiz
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String {
			trimmedValue := strings.TrimSpace(field.String())
			field.SetString(trimmedValue)
		}
	}
}

func ValidateRequestModel(s interface{}) []ValidationError {
	var errors []ValidationError
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).Interface().(string)
		fieldType := v.Type().Field(i)
		tag := fieldType.Tag.Get("validate")

		tags := strings.Split(tag, ",")

		for _, rule := range tags {
			switch {
			case rule == "required":
				if err := Required(fieldValue); err != nil {
					errors = append(errors, ValidationError{
						Field:   fieldType.Name,
						Message: err.Error(),
					})
				}

			case strings.HasPrefix(rule, "min="):
				minValue, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
				if err := Min(fieldValue, minValue); err != nil {
					errors = append(errors, ValidationError{
						Field:   fieldType.Name,
						Message: err.Error(),
					})
				}

			case strings.HasPrefix(rule, "max="):
				maxValue, _ := strconv.Atoi(strings.TrimPrefix(rule, "max="))
				if err := Max(fieldValue, maxValue); err != nil {
					errors = append(errors, ValidationError{
						Field:   fieldType.Name,
						Message: err.Error(),
					})
				}

			case rule == "email":
				if err := Email(fieldValue); err != nil {
					errors = append(errors, ValidationError{
						Field:   fieldType.Name,
						Message: err.Error(),
					})
				}
			}
		}
	}

	return errors
}

func Required(value string) error {
	if value == "" {
		return errors.New("field is required")
	}
	return nil
}

func Min(value string, minLength int) error {
	if utf8.RuneCountInString(value) < minLength {
		return errors.New("field length is less than minimum required")
	}
	return nil
}

func Max(value string, maxLength int) error {
	if utf8.RuneCountInString(value) > maxLength {
		return errors.New("field length exceeds maximum allowed")
	}
	return nil
}

func Email(value string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(value) {
		return errors.New("invalid email format")
	}
	return nil
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field '%s': %s", e.Field, e.Message)
}
