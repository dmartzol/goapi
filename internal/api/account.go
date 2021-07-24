package api

import "fmt"

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
