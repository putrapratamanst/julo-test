package utils

import (
	"julo-test/pkg"
	"julo-test/presenter"

	"github.com/gin-gonic/gin"
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
	} else {
		return &presenter.Response{
			Message: pkg.ErrFormatRequestBody.Error(),
		}
	}

	return nil
}
