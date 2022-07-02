package wallet

import (
	"julo-test/model/wallet"
	"julo-test/presenter"
)

type IWalletService interface {
	EnableWalletService(model wallet.WalletModel) *presenter.Response
	CheckWalletService(cid string) bool
	ViewWalletService(cid string) (wallet.WalletModel, *presenter.Response)
	DisableWalletService(cid string)
	DepositWalletService(data wallet.WalletModel, model wallet.DepositModel)
	WithdrawalWalletService(data wallet.WalletModel, model wallet.WithdrawalModel)
}
