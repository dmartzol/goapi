// Package proto provides helper functions related to protocol buffers
package proto

import (
	"github.com/dmartzol/goapi/goapi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// CoreAccount converts a proto account struct to goapi.Account
func CoreAccount(account *Account) (*goapi.Account, error) {
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

func AccountProto(account *goapi.Account) (*Account, error) {
	new := Account{
		Id:        account.ID.String(),
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}

	return &new, nil
}

// ToCoreAccount converts AddAccountMessage to goapim.Account
func (a *AddAccountMessage) ToCoreAccount() *goapi.Account {
	res := goapi.Account{
		FirstName: a.FirstName,
		LastName:  a.LastName,
		Email:     a.Email,
	}
	return &res
}
