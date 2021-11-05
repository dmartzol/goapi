package postgres

import (
	"github.com/dmartzol/goapi/goapi"
	"github.com/pkg/errors"
)

// AccountWithCredentials returns an account if the email and password provided match an (email,password) pair in the db
func (db *DB) AccountWithCredentials(email, password string) (*goapi.Account, error) {
	var a goapi.Account
	sqlSelect := `select * from accounts a where a.email = $1 and a.passhash = crypt($2, a.passhash)`
	err := db.Client.Get(&a, sqlSelect, email, password)
	return &a, err
}

// AddAccount insert a new account in the database
func (db *DB) AddAccount(a *goapi.Account) (*goapi.Account, error) {
	a.Model = goapi.NewModel()

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
		a.Model.ID,
		a.FirstName,
		a.LastName,
		a.Gender,
		a.Email,
		a.PassHash,
		a.CreatedTime,
		a.UpdatedTime,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert account")
	}
	return a, nil
}
