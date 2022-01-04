package helper

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateStruct(data interface{}) (error){
	validate = validator.New()
	err := validate.Struct(data)
	return err
}