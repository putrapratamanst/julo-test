package wallet

import (
	"julo-test/model/wallet"
)

type IWalletRepository interface {
	Enable(model wallet.WalletModel)
	Get(cid string) string
	Disable(model wallet.WalletModel)
	Deposit(model wallet.DepositModel)
	Withdrawal(model wallet.WithdrawalModel)
	GetDeposit(model wallet.DepositModel) string
	GetWithdrawal(model wallet.WithdrawalModel) string
}
