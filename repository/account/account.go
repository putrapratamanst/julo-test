package repository

import "cloud.google.com/go/firestore"

type AccountRepository struct {
	db *firestore.Client
}

//Account new repository
func NewRepository(db *firestore.Client) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (repository *AccountRepository) GetPlayerByName() bool {
	return false
}