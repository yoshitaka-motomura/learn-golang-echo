package utils

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ValidationErrorは各フィールドごとのバリデーションエラーを表す
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"error"`
}

// ValidationErrorsはバリデーションエラーのスライスを保持するカスタムエラー型
type ValidationErrors []ValidationError

// ErrorメソッドでValidationErrorsをerrorインターフェースとして実装
func (ve ValidationErrors) Error() string {
	bytes, _ := json.Marshal(ve)
	return string(bytes)
}

// ValidateStructは任意の構造体をバリデーションする共通関数
func ValidateStruct(s interface{}) error {
	v := validator.New()
	err := v.Struct(s)
	if err == nil {
		return nil
	}

	var errors ValidationErrors
	for _, err := range err.(validator.ValidationErrors) {
		var message string
		switch err.Tag() {
		case "required":
			message = "is required"
		case "min":
			message = "is too short"
		case "max":
			message = "is too long"
		case "oneof":
			message = "must be one of the allowed values"
		default:
			message = "is invalid"
		}
		errors = append(errors, ValidationError{
			Field:   err.Field(),
			Message: fmt.Sprintf("%s %s", err.Field(), message),
		})
	}
	return errors
}
