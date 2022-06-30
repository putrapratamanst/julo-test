package account

import (
	b64 "encoding/base64"
	repo "julo-test/interfaces/account"
	"julo-test/model/account"
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

func (s *Service) CreateUser(model *account.Init) {
	model.Token = b64.StdEncoding.EncodeToString([]byte(model.CustomerXID))
	s.repo.Insert(model)
}
