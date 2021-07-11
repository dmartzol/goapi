package postgres

import (
	"time"

	models "github.com/dmartzol/api-template/internal/model"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	ErrExpiredSession error
)

// SessionFromToken fetches a session by its token
func (db *DB) SessionFromToken(token string) (*models.Session, error) {
	var s models.Session
	sqlStatement := `SELECT * FROM sessions WHERE token = $1`
	err := db.Get(&s, sqlStatement, token)
	return &s, err
}

// CreateSession creates a new session
func (db *DB) AddSession(session *models.Session) (*models.Session, error) {
	if err := session.Validate(); err != nil {
		return nil, err
	}

	session.Model = models.NewModel()

	var s models.Session
	sqlInsert := `INSERT INTO sessions (account_id) VALUES ($1)`
	err := db.Get(&s, sqlInsert, session.AccountID)

	if err != nil {
		return nil, err
	}

	return &s, nil
}

// CreateSession creates a new session
func (db *DB) CreateSession(accountID uuid.UUID) (*models.Session, error) {
	var s models.Session
	sqlInsert := `INSERT INTO sessions (account_id) VALUES ($1) RETURNING *`
	err := db.Get(&s, sqlInsert, accountID)
	return &s, err
}

// ExpireSessionFromToken expires the session with the given token
func (db *DB) ExpireSessionFromToken(token string) (*models.Session, error) {
	var s models.Session
	sqlStatement := `UPDATE sessions SET expiration_time = current_timestamp WHERE token = $1 RETURNING *`
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
	sqlStatement := `SELECT * FROM sessions WHERE token = $1`
	tx.Get(&session, sqlStatement, token)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(err, "error fetching session from token %s", token)
	}
	if session.ExpirationTime.Before(time.Now()) {
		return nil, ErrExpiredSession
	}
	var updatedSession models.Session
	sqlStatement = `UPDATE sessions SET last_activity_time=default WHERE token = $1 RETURNING *`
	err = tx.Get(&updatedSession, sqlStatement, token)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(err, "error updating session from token %s", token)
	}
	return &updatedSession, tx.Commit()
}
