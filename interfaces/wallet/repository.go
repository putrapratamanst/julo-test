package wallet

import (
	"julo-test/model/wallet"
)

type IWalletRepository interface {
	Enable(model wallet.WalletModel)
	Get(cid string) string
}
