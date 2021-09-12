package storage

import "github.com/dmartzol/api-template/internal/model"

type Storage struct {
	database databaseInterface
}

type databaseInterface interface {
	AddAccount(a *model.Account) (*model.Account, error)
}

func New(db databaseInterface) *Storage {
	return &Storage{database: db}
}

func (s *Storage) AddAccount(a *model.Account) (*model.Account, error) {
	return s.database.AddAccount(a)
}
