package validator_util

import (
	"sync"

	"github.com/go-playground/locales/en_US"
	"github.com/go-playground/locales/id_ID"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/id"
	"github.com/samber/mo"
)

var (
	validate     *validator.Validate
	once         sync.Once
	utTranslator = ut.New(en_US.New(), id_ID.New())
)

func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()

		enTranslator, _ := utTranslator.GetTranslator("en_US")
		en.RegisterDefaultTranslations(validate, enTranslator)

		idTranslator, _ := utTranslator.GetTranslator("id_ID")
		id.RegisterDefaultTranslations(validate, idTranslator)
	})

	return validate
}

func GetTranslator(locale string) ut.Translator {
	return mo.TupleToOption(utTranslator.FindTranslator(locale)).OrElse(utTranslator.GetFallback())
}
