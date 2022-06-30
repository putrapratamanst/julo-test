package router

import (
	controller "julo-test/controller/account"
	"julo-test/service/account"

	"github.com/gin-gonic/gin"
)

func RouteAccount(v1 *gin.RouterGroup, ias *account.Service) {
	handler := controller.AccountControllerHandler(ias)
	account := v1.Group("")
	{
		account.POST("init", handler.Init)
	}
}
