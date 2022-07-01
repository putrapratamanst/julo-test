package account

import (
	b64 "encoding/base64"
	repo "julo-test/interfaces/account"
	"julo-test/model/account"
	"julo-test/presenter"
	"os"
)

//Service interface
type Service struct {
	repo repo.IAccountRepository
}

//NewService create new use case
func NewService(repo repo.IAccountRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUserService(model *account.Init) {
	if model.CustomerXID == os.Getenv("CUSTOMERXID") {
		model.Token = os.Getenv("TOKEN")
	} else {
		model.Token = b64.StdEncoding.EncodeToString([]byte(model.CustomerXID))
	}
	s.repo.Insert(model)
}

func (s *Service) GetUserService(token string) (account.Detail, *presenter.Response) {
	var data account.Detail
	detailUser, errDetailUser := s.repo.Detail(token)
	if errDetailUser != nil {
		return data, &presenter.Response{
			Message: errDetailUser.Message,
		}
	}
	data.CustomerXID = detailUser
	return data, nil
}
