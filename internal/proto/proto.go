// Package proto provides helper functions related to protocol buffers
package proto

import (
	"github.com/dmartzol/goapi/goapi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// GoapiAccount converts a proto account struct to goapi.Account
func GoapiAccount(account *Account) (*goapi.Account, error) {
	a := goapi.Account{
		FirstName: account.FirstName,
		LastName:  account.LastName,
	}
	id, err := uuid.FromBytes(account.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert byte[] to uuid")
	}
	a.ID = id
	return &a, nil
}
