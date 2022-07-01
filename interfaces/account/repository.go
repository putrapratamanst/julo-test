package account

import (
	"julo-test/model/account"
	"julo-test/presenter"
)

type IAccountRepository interface {
	Insert(model *account.Init)
	Detail(token string) (string, *presenter.Response)
}
