package ivalidate

import (
	"github.com/go-playground/validator"
)

// Validates models before do something
func IsValid(m interface{}) (bool, error) {
	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
