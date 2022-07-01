package wallet

import (
	"julo-test/pkg/utils"
	"julo-test/service/wallet"
)

var (
	ValidateRequest = utils.ValidateRequest
)

type WalletController struct {
	iws *wallet.Service
}

func WalletControllerHandler(iws *wallet.Service) *WalletController {
	handler := &WalletController{
		iws: iws,
	}
	return handler
}
