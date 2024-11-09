package http_controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	domain_errors "github.com/kbiits/dealls-take-home-test/domain/errors"
	validator_util "github.com/kbiits/dealls-take-home-test/utils/validator"
)

func HandleError(c *gin.Context, err error) {
	switch {
	case errors.As(err, &domain_errors.DomainError{}):
		domainErr := err.(domain_errors.DomainError)

		switch {
		case domainErr.IsResourceNotFound():
			c.JSON(404, NewErrorStringResponse(domainErr.Error()))
		case domainErr.IsValidationError():
			c.JSON(400, NewErrorStringResponse(domainErr.Error()))
		default:
			c.JSON(500, NewErrorStringResponse(domainErr.Error()))
		}

		return
	case errors.As(err, &validator.ValidationErrors{}):
		validationErr := err.(validator.ValidationErrors)
		handleValidationError(c, validationErr)
		return

	default:
		c.JSON(500, NewErrorResponse(err))
		return
	}

}

func handleValidationError(c *gin.Context, err validator.ValidationErrors) {
	firstErr := err[0]
	stringErr := firstErr.Translate(validator_util.GetTranslator("id_ID"))
	c.JSON(400, NewErrorStringResponse(stringErr))
}
