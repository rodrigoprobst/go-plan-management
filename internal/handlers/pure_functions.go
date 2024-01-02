package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rodrigoprobst/go-plan-management/pkg/validation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func handleBindError(err error, c *gin.Context) {
	msg := err.Error()
	switch errType := err.(type) {
	case *json.UnmarshalTypeError:
		msg = errType.Field + " type error"
	case validator.ValidationErrors:
		errs := err.(validator.ValidationErrors)
		response := validation.ValidationErrorsToMapResponse(errs)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": response})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"errors": gin.H{"message": msg}})
}
