package utils

import (
	"fmt"
	"julo-test/pkg"
	"julo-test/presenter"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/id"
	UNIV_TRANS "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ID_TRANS "github.com/go-playground/validator/v10/translations/id"
)

func ValidateRequest(bindType int, ctx *gin.Context, input interface{}) *presenter.Response {
	//check request body
	if bindType == pkg.BIND_TYPE_JSON {
		if errBind := ctx.ShouldBindJSON(&input); errBind != nil {
			return &presenter.Response{
				Message: pkg.ErrFormatRequestBody.Error(),
			}
		}
	} else if bindType == pkg.BIND_TYPE_PARAM {
		if errBind := ctx.ShouldBindQuery(input); errBind != nil {
			return &presenter.Response{
				Message: pkg.ErrFormatRequestBody.Error(),
			}
		}
	} else if bindType == pkg.BIND_TYPE_FORM {
		//nanti aja erornya dihandle
		if errBind := ctx.ShouldBind(&input); errBind != nil {
			return &presenter.Response{
				Message: pkg.ErrFormatRequestBody.Error(),
			}
		}
	} else {
		return &presenter.Response{
			Message: pkg.ErrFormatRequestBody.Error(),
		}
	}

	//validate request body
	validate := validator.New()
	uni := UNIV_TRANS.New(id.New())
	trans, _ := uni.GetTranslator("id")

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
		if name == "" {
			return ""
		}

		return name
	})

	//Verifier registration translator
	errTranslation := ID_TRANS.RegisterDefaultTranslations(validate, trans)
	if errTranslation != nil {
		return &presenter.Response{
			Code:    http.StatusBadRequest,
			Message: errTranslation.Error(),
		}
	}

	errTranslation = validate.Struct(input)
	msgError := ""

	if errTranslation != nil {
		for _, e := range errTranslation.(validator.ValidationErrors) {
			fmt.Println(e)
			translatedErr := fmt.Errorf(e.Translate(trans))
			msgError = msgError + translatedErr.Error() + ". "
		}

		return &presenter.Response{
			Code:    http.StatusBadRequest,
			Message: msgError,
		}
	}

	return nil
}
