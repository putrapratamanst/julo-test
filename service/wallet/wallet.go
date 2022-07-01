package wallet

import (
	"encoding/json"
	repo "julo-test/interfaces/wallet"
	"julo-test/model/wallet"
	"julo-test/pkg"
)

//Service interface
type Service struct {
	repo repo.IWalletRepository
}

//NewService create new use case
func NewService(repo repo.IWalletRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) EnableWalletService(model wallet.WalletModel) {
	s.repo.Enable(model)
}

func (s *Service) CheckWalletService(cid string) (bool, int) {
	check := s.repo.Get(cid)
	if check == "" { //new data
		return true, 0
	} else {
		var data wallet.WalletModel
		json.Unmarshal([]byte(check), &data)
		if data.Status == pkg.WALLET_ENABLED {
			return false, data.Balance
		} else {
			return true, data.Balance
		}
	}
}
