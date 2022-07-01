package wallet

import "julo-test/service/wallet"

type WalletController struct {
	iws *wallet.Service
}

func WalletControllerHandler(iws *wallet.Service) *WalletController {
	handler := &WalletController{
		iws: iws,
	}
	return handler
}
