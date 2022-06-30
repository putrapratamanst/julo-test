package account

import (
	model "julo-test/model/account"
	"julo-test/presenter"
)

type IAccountService interface {
	Init(input *model.Init) *presenter.Response
}
