package model

import (
	"fmt"
	"time"
)

// Session represents an account session
type Session struct {
	*Model
	AccountID        int64     `db:"account_id"`
	Token            string    `db:"token"`
	LastActivityTime time.Time `db:"last_activity_time"`
	ExpirationTime   time.Time `db:"expiration_time"`
}

func (s *Session) Validate() error {
	if s.ExpirationTime.After(time.Now()) {
		return fmt.Errorf("session expired")
	}
	return nil
}

type SessionView struct {
	AccountID        int64
	LastActivityTime time.Time
	ExpirationTime   time.Time
}

func (s Session) View(options map[string]bool) SessionView {
	view := SessionView{
		AccountID:        s.AccountID,
		LastActivityTime: s.LastActivityTime,
		ExpirationTime:   s.ExpirationTime,
	}
	return view
}

type LoginCredentials struct {
	Email    string
	Password string
}
