package account

import (
	"julo-test/model/account"
)

type IAccountRepository interface {
	Insert(model *account.Init)
}
