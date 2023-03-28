package goapi

import (
	"strings"

	"github.com/pkg/errors"
)

type Accounts []*Account

// Account represents a user account
type Account struct {
	*Model
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Gender    *string
	PassHash  string `db:"pass_hash"`
	Email     string
}

type RegisterRequest struct {
	FirstName   string
	LastName    string
	Gender      *string
	PhoneNumber *string
	Email       string
	Password    string
}

func (r RegisterRequest) Validate() error {
	if r.FirstName == "" {
		return errors.New("first name is required")
	}
	if r.LastName == "" {
		return errors.New("last name is required")
	}
	if r.Email == "" {
		return errors.New("email is required")
	}
	ok := validEmail(r.Email)
	if !ok {
		return errors.New("invalid email")
	}
	if len(r.Password) < 6 {
		return errors.New("password too short")
	}
	if r.Gender != nil && *r.Gender != "" && *r.Gender != "M" && *r.Gender != "F" {
		return errors.New("Gender value not implemented")
	}
	return nil
}

func validEmail(email string) bool {
	if !strings.Contains(email, "@") {
		return false
	}
	if !strings.Contains(email, ".") {
		return false
	}
	return true
}

type ResetPasswordRequest struct {
	Email string
}

type ConfirmEmailRequest struct {
	ConfirmationKey string
}
