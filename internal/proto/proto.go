// Package proto provides helper functions related to protocol buffers
package proto

import (
	"github.com/dmartzol/goapi/goapi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// GoapiAccount converts a proto account struct to goapi.Account
func GoapiAccount(account *Account) (*goapi.Account, error) {
	new := goapi.Account{
		Model:     &goapi.Model{},
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}

	new.CreatedTime = account.CreatedTime.AsTime()
	if account.UpdatedTime != nil {
		t := account.UpdatedTime.AsTime()
		new.UpdatedTime = &t
	}

	id, err := uuid.Parse(account.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert byte[] to uuid")
	}
	new.ID = id

	return &new, nil
}
