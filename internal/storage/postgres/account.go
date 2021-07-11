package postgres

import "github.com/dmartzol/api-template/internal/model"

// AccountWithCredentials returns an account if the email and password provided match an (email,password) pair in the db
func (db *DB) AccountWithCredentials(email, password string) (*model.Account, error) {
	var a model.Account
	sqlStatement := `select * from accounts a where a.email = $1 and a.passhash = crypt($2, a.passhash)`
	err := db.Get(&a, sqlStatement, email, password)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
