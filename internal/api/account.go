package api

import (
	"fmt"

	"github.com/google/uuid"
)

// AccountView is the restricted response body of Account
// see: https://stackoverflow.com/questions/46427723/golang-elegant-way-to-omit-a-json-property-from-being-serialized
type Account struct {
	ID                         uuid.UUID `json:"ID"`
	FirstName, LastName, Email string
	DOB                        string
	Gender                     string
}

type CreateAccountRequest struct {
	FirstName string
	LastName  string
	Email     string
}

func (c *CreateAccountRequest) Validate() error {
	if c.FirstName == "" {
		return fmt.Errorf("empty first name error")
	}
	if c.LastName == "" {
		return fmt.Errorf("empty last name error")
	}
	if c.Email == "" {
		return fmt.Errorf("empty email error")
	}
	return nil
}
