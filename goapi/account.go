package goapi

import (
	"strings"
	"time"

	"github.com/dmartzol/goapi/internal/api"
	"github.com/dmartzol/goapi/pkg/timeutils"
	"github.com/pkg/errors"
)

type Accounts []*Account

// Account represents a user account
type Account struct {
	*Model
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	DOB       time.Time
	Gender    *string
	PassHash  string `db:"pass_hash"`
	Email     string
}

// View returns the Account struct restricted to those fields allowed in options
// see: https://stackoverflow.com/questions/46427723/golang-elegant-way-to-omit-a-json-property-from-being-serialized
func (a Account) View(options map[string]bool) api.Account {
	view := api.Account{
		ID:        a.ID,
		FirstName: a.FirstName,
		LastName:  a.LastName,
		DOB:       a.DOB.Format(timeutils.LayoutISODay),
		Email:     a.Email,
	}
	if a.Gender != nil {
		view.Gender = *a.Gender
	}
	return view
}

func (accs Accounts) Views(options map[string]bool) []api.Account {
	var l []api.Account
	for _, a := range accs {
		l = append(l, a.View(options))
	}
	return l
}

type RegisterRequest struct {
	FirstName   string
	LastName    string
	DOB         string
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
