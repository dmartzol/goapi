package postgres

import (
	"github.com/dmartzol/goapi/goapi"
	"github.com/pkg/errors"
)

type Account struct {
	*goapi.Account
}

func (a *Account) Validate() error {
	if a.FirstName == "" {
		return errors.Errorf("invalid name")
	}
	if a.LastName == "" {
		return errors.Errorf("invalid last name")
	}
	if a.Email == "" {
		return errors.Errorf("empty email")
	}
	return nil
}

func (a *Account) Build() *goapi.Account {
	return a.Account
}

// AccountWithCredentials returns an account if the email and password provided match an (email,password) pair in the db
func (db *DB) AccountWithCredentials(email, password string) (*goapi.Account, error) {
	var a goapi.Account
	sqlSelect := `select * from accounts a where a.email = $1 and a.passhash = crypt($2, a.passhash)`
	err := db.Client.Get(&a, sqlSelect, email, password)
	return &a, err
}

// AddAccount insert a new account in the database
func (db *DB) AddAccount(a *goapi.Account) (*goapi.Account, error) {
	dbAccount := &Account{
		Account: a,
	}
	if err := dbAccount.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation failed")
	}

	dbAccount.Model = goapi.NewModel()

	sqlInsert := `
	insert into accounts (
		id,
		first_name,
		last_name,
		gender,
		email,
		pass_hash,
		created_time,
		updated_time)
		values
		($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := db.Client.Exec(
		sqlInsert,
		dbAccount.Model.ID,
		dbAccount.FirstName,
		dbAccount.LastName,
		dbAccount.Gender,
		dbAccount.Email,
		dbAccount.PassHash,
		dbAccount.CreatedTime,
		dbAccount.UpdatedTime,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert account")
	}
	return dbAccount.Build(), nil
}
