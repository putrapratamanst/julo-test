package wallet

import (
	"julo-test/model/wallet"
	"julo-test/presenter"
)

type IWalletService interface {
	EnableWalletService(model wallet.WalletModel) *presenter.Response
	CheckWalletService(cid string) bool
}
