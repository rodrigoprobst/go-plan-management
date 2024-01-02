package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitializeValidatorConfigs() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func ValidationErrorsToMapResponse(err validator.ValidationErrors) map[string]string {
	list := make(map[string]string)
	for _, e := range err {
		list[e.Field()] = validationErrorToText(e)
	}
	return list
}

func validationErrorToText(e validator.FieldError) string {
	word := split(e.Field())

	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", word)
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", word, e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", word, e.Param())
	case "email":
		return "Invalid email format"
	case "len":
		return fmt.Sprintf("%s must be %s characters long", word, e.Param())
	case "numeric":
		return fmt.Sprintf("%s must be numeric", word)
	case "alphanum":
		return fmt.Sprintf("%s must be contain only alpha numeric characters", word)
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", word, strings.Join(strings.Split(e.Param(), " "), ","))
	case "password":
		return fmt.Sprintf("%s must be contain at least 1 UpperCase letter, 1 LowerCase letter, 1 number and 1 symbol.", word)
	}
	return fmt.Sprintf("%s is not valid", word)
}

func split(src string) string {
	split := strings.Split(src, "_")
	justString := strings.Join(split, " ")
	return justString
}
