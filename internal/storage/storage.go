package storage

import "github.com/dmartzol/goapi/goapi"

type MacroStorage struct {
	database storageInterface
}

type storageInterface interface {
	AddAccount(a *goapi.Account) (*goapi.Account, error)
}

func New(db storageInterface) *MacroStorage {
	return &MacroStorage{database: db}
}

func (s *MacroStorage) AddAccount(a *goapi.Account) (*goapi.Account, error) {
	a, err := s.database.AddAccount(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
