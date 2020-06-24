package validation

import (
	"fmt"

	"github.com/arifseft/go-actions/middlewares/exception"
	"github.com/go-playground/validator/v10"
)

func Validate(schema interface{}) {
	validate := validator.New()

	if err := validate.Struct(schema); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			exception.BadRequest(fmt.Sprint(err), "INVALID_BODY")
			return
		}
		for _, err := range err.(validator.ValidationErrors) {
			exception.BadRequest(fmt.Sprint(err), "INVALID_BODY")
		}
		return
	}
}

