package account

import (
	"julo-test/model/account"
	"julo-test/pkg"
	"julo-test/presenter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AccountController) Init(ctx *gin.Context) {
	var input account.Init
	err := ValidateRequest(pkg.BIND_TYPE_JSON, ctx, &input)
	if err != nil {
		result := presenter.Response{
			Status:  pkg.HTTP_STATUS_FAIL,
			Message: err.Message,
			Data:    nil,
		}
		ctx.JSON(err.Code, result)
		return
	}

	handler.ias.CreateUser(&input)
	result := presenter.Response{
		Status: pkg.HTTP_STATUS_SUCCESS,
		Data: presenter.InitResponse{
			Token: input.Token,
		},
	}
	ctx.JSON(http.StatusOK, result)
}
