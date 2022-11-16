package validate

import (
	"github.com/go-playground/validator/v10"
)

// Validator is the default validator
var validate = validator.New()

type ErrorInfo struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func Validate(req interface{}) []*ErrorInfo {
	var errors []*ErrorInfo
	err := validate.Struct(req)
	if err != nil {

		for _, inputErr := range err.(validator.ValidationErrors) {
			errInfo := ErrorInfo{
				FailedField: inputErr.StructNamespace(),
				Tag:         inputErr.Tag(),
				Value:       inputErr.Param(),
			}
			errors = append(errors, &errInfo)
		}
	}

	return errors
}
