package helper

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/backend-magang/eniqilo-store/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func UseCustomValidatorHandler(c *echo.Echo) {
	c.Validator = &CustomValidator{validator: validator.New()}
	c.HTTPErrorHandler = func(err error, c echo.Context) {
		var MessageValidation []string
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required", "required_with", "required_without", "required_unless":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s is required",
						err.Field()))
				case "email":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s is not valid email",
						err.Field()))
				case "gte":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param()))
				case "lte":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param()))
				case "numeric|eq=*":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value must be numeric or *",
						err.Field()))
				case "oneof":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value not one of %s",
						err.Field(), strings.ReplaceAll(err.Param(), " ", "/")))
				}
			}
			_ = WriteResponse(c, models.StandardResponseReq{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Data:    MessageValidation,
				Error:   nil,
			})
		}
	}

}

func IsValidSlug(s string) bool {
	regex, _ := regexp.Compile(`^[a-z0-9-]+$`)
	return regex.MatchString(s)
}

func IsNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func IsInArray(str string, arr []string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}
