package cmvalidator

import (
	"bytes"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

const CustomMessageTag = "customMessage"

// CMValidate is a custom validator that extends the go-playground/validator/v10 package.
type CMValidate struct {
	*validator.Validate
}

// New creates a new instance of the CMValidate struct.
func New() *CMValidate {
	return &CMValidate{validator.New()}
}

// Struct validates a struct and returns an error if the struct is invalid.
func (cv *CMValidate) Struct(s any) (err error) {
	return cv.StructCtx(context.Background(), s)
}

// StructCtx validates a struct with a context and returns an error if the struct is invalid.
func (cv *CMValidate) StructCtx(ctx context.Context, s any) (err error) {
	err = cv.Validate.StructCtx(ctx, s)

	if err == nil {
		return
	}

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return
	}

	return cnvToCMValidateErrors(validationErrors, s)
}

func cnvToCMValidateErrors(err validator.ValidationErrors, s any) CFValidateErrors {
	cmvalidateErrors := make(CFValidateErrors, 0, len(err))
	for _, validationError := range err {
		cmvalidateErrors = append(cmvalidateErrors, cnvToCMFieldError(validationError, s))
	}

	return cmvalidateErrors

}

func cnvToCMFieldError(err validator.FieldError, s any) CMFieldError {

	typ := reflect.TypeOf(s)
	fieldName := err.StructField()

	field, ok := typ.FieldByName(fieldName)
	if !ok {
		return cmFieldError{
			FieldError:    err,
			customMessage: "",
		}
	}

	customMessage := field.Tag.Get(CustomMessageTag)

	return cmFieldError{
		FieldError:    err,
		customMessage: customMessage,
	}
}

// CFValidateErrors is a slice of CMFieldError.
// get the custom error message by calling the CustomMessage method.
type CFValidateErrors []CMFieldError

// Error is the same as ValidationErrors and is intended for development and debugging.
func (cfe CFValidateErrors) Error() string {
	buff := bytes.NewBufferString("")

	for i := 0; i < len(cfe); i++ {
		buff.WriteString(cfe[i].Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

type CMFieldError interface {
	validator.FieldError
	// CustomMessage returns the value of the customMessage tag.
	// If no custom error message is found, an empty string is returned.
	CustomMessage() string
}

type cmFieldError struct {
	validator.FieldError
	customMessage string
}

// CustomMessage returns the value of the customMessage tag.
func (cfe cmFieldError) CustomMessage() string {
	return cfe.customMessage
}
