package server

import (
	"errors"
	"log"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

type CustomValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

func NewCustomValidator() *CustomValidator {
	validate := validator.New()
	uni := ut.New(en.New(), zh.New())
	trans, _ := uni.GetTranslator("zh")
	if err := zhTrans.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatalf("validator register translations failed: %v", err)
	}
	return &CustomValidator{validator: validate, trans: trans}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				return errors.New(e.Translate(cv.trans))
			}
		}
		return err
	}
	return nil
}
