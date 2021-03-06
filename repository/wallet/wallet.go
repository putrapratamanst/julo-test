package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	"julo-test/infrastructure"
	"julo-test/model/wallet"
)

type WalletRepository struct {
	rc *infrastructure.RedisCache
}

//Wallet new repository
func NewRepository(rc *infrastructure.RedisCache) *WalletRepository {
	return &WalletRepository{
		rc: rc,
	}
}

func (repository *WalletRepository) Enable(model wallet.WalletModel) {
	context := context.Background()
	setValue, _ := json.Marshal(model)
	repository.rc.Client.Set(context, "wallet:"+model.OwnedBy, setValue, 0)
}

func (repository *WalletRepository) Get(cid string) string {
	context := context.Background()
	dataKey := fmt.Sprintf("wallet:%s", cid)
	data, _ := repository.rc.Client.Get(context, dataKey).Result()
	return data
}

func (repository *WalletRepository) Disable(model wallet.WalletModel) {
	context := context.Background()
	setValue, _ := json.Marshal(model)
	repository.rc.Client.Set(context, "wallet:"+model.OwnedBy, setValue, 0)
}

func (repository *WalletRepository) Deposit(model wallet.DepositModel) {
	context := context.Background()
	setValue, _ := json.Marshal(model)
	repository.rc.Client.Set(context, model.DepositedBy+":deposit:"+model.ReferenceID, setValue, 0)
}

func (repository *WalletRepository) Withdrawal(model wallet.WithdrawalModel) {
	context := context.Background()
	setValue, _ := json.Marshal(model)
	repository.rc.Client.Set(context, model.WithdrawnBy+":withdrawal:"+model.ReferenceID, setValue, 0)
}

func (repository *WalletRepository) GetDeposit(model wallet.DepositModel) string {
	context := context.Background()
	dataKey := fmt.Sprintf("%s:deposit:%s", model.DepositedBy, model.ReferenceID)
	data, _ := repository.rc.Client.Get(context, dataKey).Result()
	return data
}

func (repository *WalletRepository) GetWithdrawal(model wallet.WithdrawalModel) string {
	context := context.Background()
	dataKey := fmt.Sprintf("%s:withdrawal:%s", model.WithdrawnBy, model.ReferenceID)
	data, _ := repository.rc.Client.Get(context, dataKey).Result()
	return data
}
