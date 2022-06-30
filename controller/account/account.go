package account

import (
	"julo-test/model/account"
	"julo-test/pkg"
	"julo-test/pkg/response"
	"julo-test/presenter"

	"github.com/gin-gonic/gin"
)

func (handler *AccountController) Init(ctx *gin.Context) {
	var input account.Init
	err := ValidateRequest(pkg.BIND_TYPE_PARAM, ctx, &input)
	if err != nil {
		result := presenter.Response{
			Code:    err.Code,
			Message: err.Message,
		}
		response.Response(ctx, &result)
		return
	}

	result := presenter.Response{
		Status: pkg.HTTP_STATUS_SUCCESS,
		Data:   nil,
	}
	response.Response(ctx, &result)

}
