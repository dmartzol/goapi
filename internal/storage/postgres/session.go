package postgres

import (
	"time"

	models "github.com/dmartzol/api-template/internal/model"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	ErrExpiredResource error
)

// SessionFromToken fetches a session by its token
func (db *DB) SessionFromToken(token string) (*models.Session, error) {
	var s models.Session
	sqlStatement := `select * from sessions where token = $1`
	err := db.Get(&s, sqlStatement, token)
	return &s, err
}

// CreateSession creates a new session
func (db *DB) CreateSession(accountID uuid.UUID) (*models.Session, error) {
	var s models.Session
	sqlInsert := `insert into sessions (account_id) values ($1) returning *`
	err := db.Get(&s, sqlInsert, accountID)
	return &s, err
}

// ExpireSessionFromToken expires the session with the given token
func (db *DB) ExpireSessionFromToken(token string) (*models.Session, error) {
	var s models.Session
	sqlStatement := `update sessions set expiration_time = current_timestamp where token = $1 returning *`
	err := db.Get(&s, sqlStatement, token)
	return &s, err
}

// UpdateSession updates a session in the db with the current timestamp
func (db *DB) UpdateSession(token string) (*models.Session, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	var session models.Session
	sqlStatement := `select * from sessions where token = $1`
	tx.Get(&session, sqlStatement, token)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(err, "error fetching session from token %s", token)
	}
	if session.ExpirationTime.Before(time.Now()) {
		return nil, ErrExpiredResource
	}
	var updatedSession models.Session
	sqlStatement = `update sessions set last_activity_time=default where token = $1 returning *`
	err = tx.Get(&updatedSession, sqlStatement, token)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(err, "error updating session from token %s", token)
	}
	return &updatedSession, tx.Commit()
}
