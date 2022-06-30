package account

import (
	"julo-test/pkg/utils"
	"julo-test/service/account"
)

var (
	ValidateRequest = utils.ValidateRequest
)

type AccountController struct {
	ias *account.Service
}

func AccountControllerHandler(ias *account.Service) *AccountController {
	handler := &AccountController{
		ias: ias,
	}
	return handler
}
