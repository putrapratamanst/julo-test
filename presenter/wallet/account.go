package presenter

import "julo-test/model/wallet"

type EnableResponse struct {
	Wallet wallet.WalletModel `json:"wallet"`
}

type DepositResponse struct {
	Deposit wallet.DepositModel `json:"deposit"`
}

type WithdrawalResponse struct {
	Withdrawal wallet.WithdrawalModel `json:"withdrawal"`
}
