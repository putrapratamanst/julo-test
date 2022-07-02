package wallet

import (
	"encoding/json"
	repo "julo-test/interfaces/wallet"
	"julo-test/model/wallet"
	"julo-test/pkg"
	"julo-test/presenter"
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

func (s *Service) ViewWalletService(cid string) (wallet.WalletModel, *presenter.Response) {
	var data wallet.WalletModel
	check := s.repo.Get(cid)
	if check == "" {
		return data, &presenter.Response{
			Message: pkg.ErrDataNotFound.Error(),
		}
	}

	json.Unmarshal([]byte(check), &data)
	return data, nil
}

func (s *Service) DisableWalletService(model wallet.WalletModel) {
	s.repo.Disable(model)
}

func (s *Service) DepositWalletService(data wallet.WalletModel, model wallet.DepositModel) {
	s.repo.Enable(data)
	s.repo.Deposit(model)
}

func (s *Service) WithdrawalWalletService(data wallet.WalletModel, model wallet.WithdrawalModel) {
	s.repo.Enable(data)
	s.repo.Withdrawal(model)
}

func (s *Service) GetDepositService(model wallet.DepositModel) (wallet.DepositModel, *presenter.Response) {
	var data wallet.DepositModel
	detail := s.repo.GetDeposit(model)
	if detail == "" {
		return data, &presenter.Response{
			Message: pkg.ErrDataNotFound.Error(),
		}
	}

	json.Unmarshal([]byte(detail), &data)
	return data, nil
}

func (s *Service) GetWithdrawalService(model wallet.WithdrawalModel) (wallet.WithdrawalModel, *presenter.Response) {
	var data wallet.WithdrawalModel
	detail := s.repo.GetWithdrawal(model)
	if detail == "" {
		return data, &presenter.Response{
			Message: pkg.ErrDataNotFound.Error(),
		}
	}

	json.Unmarshal([]byte(detail), &data)
	return data, nil
}
