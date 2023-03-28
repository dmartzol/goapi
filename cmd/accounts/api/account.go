package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dmartzol/goapi/goapi"
	"github.com/dmartzol/goapi/internal/proto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AccountView is the restricted response body of Account
// see: https://stackoverflow.com/questions/46427723/golang-elegant-way-to-omit-a-json-property-from-being-serialized
type Account struct {
	ID                         uuid.UUID `json:"ID"`
	FirstName, LastName, Email string
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

// ToAPIAccount returns the Account struct restricted to those fields allowed in options
// see: https://stackoverflow.com/questions/46427723/golang-elegant-way-to-omit-a-json-property-from-being-serialized
func ToAPIAccount(internalAccount *goapi.Account, options map[string]bool) Account {
	apiAccount := Account{
		ID:        internalAccount.ID,
		FirstName: internalAccount.FirstName,
		LastName:  internalAccount.LastName,
		Email:     internalAccount.Email,
	}
	if internalAccount.Gender != nil {
		apiAccount.Gender = *internalAccount.Gender
	}
	return apiAccount
}

func Views(accs goapi.Accounts, options map[string]bool) []Account {
	var l []Account
	for _, a := range accs {
		l = append(l, ToAPIAccount(a, options))
	}
	return l
}

func (h *Handler) createAccount(c *gin.Context) {
	var req CreateAccountRequest
	err := h.Unmarshal(c, &req)
	if err != nil {
		h.Errorw("could not unmarshal", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		h.Errorw("failed to validate create account request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addAccountMessage := proto.AddAccountMessage{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pbAccount, err := h.Accounts.AddAccount(ctx, &addAccountMessage)
	if err != nil {
		h.Errorw("failed to create account", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	a, err := proto.CoreAccount(pbAccount)
	if err != nil {
		h.Errorw("failed to marshall account", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	vOpts := make(map[string]bool)

	c.JSON(http.StatusOK, ToAPIAccount(a, vOpts))
}
