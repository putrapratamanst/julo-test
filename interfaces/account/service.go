package account

import (
	model "julo-test/model/account"
	"julo-test/presenter"
)

type IAccountService interface {
	CreateUser(input model.Init) *presenter.Response
}
