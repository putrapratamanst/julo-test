package account
import (
	repo "julo-test/interfaces/account"
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
