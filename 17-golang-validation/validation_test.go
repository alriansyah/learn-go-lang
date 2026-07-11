package golang_validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New() // inti dari package validator, dan merupakan Instance dari struct Validate
	if validate == nil {
		t.Error("Validate is nil")
	}
}
