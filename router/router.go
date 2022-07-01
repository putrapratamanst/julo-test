package router

import (
	accountController "julo-test/controller/account"
	walletController "julo-test/controller/wallet"
	"julo-test/middleware"
	"julo-test/service/account"
	"julo-test/service/wallet"

	"github.com/gin-gonic/gin"
)

func RouteAccount(v1 *gin.RouterGroup, ias *account.Service) {
	handler := accountController.AccountControllerHandler(ias)
	account := v1.Group("")
	{
		account.POST("init", handler.Init)
	}
}

func RouteWallet(v1 *gin.RouterGroup, iws *wallet.Service, ias *account.Service) {
	handler := walletController.WalletControllerHandler(iws)
	wallet := v1.Group("/wallet")
	{
		wallet.POST("", middleware.AuthUser(ias), handler.Enable)
		wallet.GET("", middleware.AuthUser(ias), handler.View)
		wallet.PATCH("", middleware.AuthUser(ias), handler.Disable)
	}
}
