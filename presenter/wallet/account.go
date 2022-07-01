package presenter

import "julo-test/model/wallet"

type EnableResponse struct {
	Wallet wallet.WalletModel `json:"wallet"`
}
