package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	v     *validator.Validate
	trans ut.Translator
}

func New() *Validator {
	v := validator.New()
	english := en.New()
	un := ut.New(english, english)
	trans, _ := un.GetTranslator("en")

	return &Validator{
		v:     v,
		trans: trans,
	}
}

func (v Validator) Validate(st any) error {
	err := v.v.Struct(st)
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errTranslations := ve.Translate(v.trans)
		errMsg := ""
		for k, v := range errTranslations {
			errMsg += fmt.Sprintf("%s: %s\n", k, v)
		}
		err = errors.New(errMsg)
	}

	return err
}
