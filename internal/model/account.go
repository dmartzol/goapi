package model

import (
	"strings"
	"time"

	pb "github.com/dmartzol/api-template/internal/protos"
	"github.com/dmartzol/api-template/pkg/timeutils"
	"github.com/google/uuid"
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

// AccountView is the restricted response body of Account
// see: https://stackoverflow.com/questions/46427723/golang-elegant-way-to-omit-a-json-property-from-being-serialized
type AccountView struct {
	ID                         uuid.UUID `json:"ID"`
	FirstName, LastName, Email string
	DOB                        string
	Gender                     string
}

// View returns the Account struct restricted to those fields allowed in options
// see: https://stackoverflow.com/questions/46427723/golang-elegant-way-to-omit-a-json-property-from-being-serialized
func (a Account) View(options map[string]bool) AccountView {
	view := AccountView{
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

func (accs Accounts) Views(options map[string]bool) []AccountView {
	var l []AccountView
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

func MarshallAccount(accountMessage *pb.AccountMessage) (*Account, error) {
	a := Account{
		FirstName: accountMessage.FirstName,
		LastName:  accountMessage.LastName,
	}
	id, err := uuid.FromBytes(accountMessage.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert byte[] to uuid")
	}
	a.ID = id
	return &a, nil
}
