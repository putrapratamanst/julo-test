package account

import (
	model "julo-test/model/account"
	"julo-test/presenter"
)

type IAccountService interface {
	CreateUserService(input model.Init) *presenter.Response
	GetUserService(token string) (model.Detail, *presenter.Response)
}
