package repository

import (
	"context"
	"fmt"
	"julo-test/infrastructure"
	"julo-test/model/account"
	"julo-test/pkg"
	"julo-test/presenter"

	"github.com/go-redis/redis"
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
	repository.rc.Client.Set(context, "account:"+model.Token, model.CustomerXID, 0)
}

func (repository *AccountRepository) Detail(token string) (string, *presenter.Response) {
	context := context.Background()
	key := fmt.Sprintf("account:%s", token)
	data, err := repository.rc.Client.Get(context, key).Result()
	if err != nil && err != redis.Nil {
		return "", &presenter.Response{
			Message: pkg.ErrGetDataRedis.Error(),
		}
	}
	return data, nil

}
