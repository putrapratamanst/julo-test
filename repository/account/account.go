package repository

import (
	"context"
	"julo-test/infrastructure"
	"julo-test/model/account"
)

type AccountRepository struct {
	rc *infrastructure.RedisCache
}

//Account new repository
func NewRepository(rc *infrastructure.RedisCache) *AccountRepository {
	return &AccountRepository{
		rc: rc,
	}
}

func (repository *AccountRepository) Insert(model *account.Init) {
	context := context.Background()
	repository.rc.Client.Set(context, "account:"+model.CustomerXID, model.Token, 0)
}
