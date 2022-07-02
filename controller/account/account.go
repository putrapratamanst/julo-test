package account

import (
	"julo-test/model/account"
	"julo-test/pkg"
	"julo-test/presenter"
	accountResponse "julo-test/presenter/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AccountController) Init(ctx *gin.Context) {
	cid := ctx.PostForm("customer_xid")
	var input account.Init
	input.CustomerXID = cid
	err := ValidateRequest(pkg.BIND_TYPE_PARAM, ctx, &input)
	if err != nil {
		result := presenter.Response{
			Status:  pkg.HTTP_STATUS_FAIL,
			Message: err.Message,
			Data:    nil,
		}
		ctx.JSON(err.Code, result)
		return
	}

	handler.ias.CreateUserService(&input)
	result := presenter.Response{
		Status: pkg.HTTP_STATUS_SUCCESS,
		Data: accountResponse.InitResponse{
			Token: input.Token,
		},
	}
	ctx.JSON(http.StatusCreated, result)
}
