package storage

import "github.com/dmartzol/goapi/goapi"

type Storage struct {
	database databaseInterface
}

type databaseInterface interface {
	AddAccount(a *goapi.Account) (*goapi.Account, error)
}

func New(db databaseInterface) *Storage {
	return &Storage{database: db}
}

func (s *Storage) AddAccount(a *goapi.Account) (*goapi.Account, error) {
	return s.database.AddAccount(a)
}
