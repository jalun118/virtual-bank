package helpers

import (
	"regexp"
	"virtual-bank/internal/constants"

	"github.com/go-playground/validator/v10"
)

const (
	alphaNumericSpecialRegexString = "^[a-zA-Z0-9_-]+$"
	numericRegexString             = "^[0-9]+$"
	notAllowCharRegexString        = `[^\\/:*?"<>|]`
)

func ValidationAlpaNumber(text string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString(text)
}

func ValidationNotAllowChar(text string) bool {
	return regexp.MustCompile(notAllowCharRegexString).MatchString(text)
}

func ValidationAlpaNumberSpecial(text string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(text)
}

func ValidationSort(text string) bool {
	return regexp.MustCompile("^(ASC|asc|DESC|desc)$").MatchString(text)
}

func alpaNumericSpecialChar(fl validator.FieldLevel) bool {
	return regexp.MustCompile(alphaNumericSpecialRegexString).MatchString(fl.Field().String())
}

func notAllowChar(fl validator.FieldLevel) bool {
	return regexp.MustCompile(notAllowCharRegexString).MatchString(fl.Field().String())
}
func numericString(fl validator.FieldLevel) bool {
	return regexp.MustCompile(numericRegexString).MatchString(fl.Field().String())
}

type Validation struct {
	validator *validator.Validate
}

func NewValidation(v *validator.Validate) *Validation {
	v.RegisterValidation("alphanumspecial", alpaNumericSpecialChar)
	v.RegisterValidation("notallowchar", notAllowChar)
	v.RegisterValidation("numstring", numericString)
	return &Validation{validator: v}
}

type errorType struct {
	Field   string
	Tag     string
	MsgCode int
}

func (v *Validation) Struct(i interface{}) []errorType {
	err := v.validator.Struct(i)
	if err != nil {
		var errors []errorType
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			msgCode := 0

			if val, ok := constants.ValidatorMsg[err.Tag()]; ok {
				msgCode = val
			}

			errorIndividual := errorType{
				Field:   err.Field(),
				Tag:     err.Tag(),
				MsgCode: msgCode,
			}
			errors = append(errors, errorIndividual)
		}
		return errors
	}
	return nil
}

func FilterTagCode(errorData []errorType, fieldFilter string) int {
	for _, v := range errorData {
		if v.Field == fieldFilter {
			return v.MsgCode
		}
	}
	return 0
}
