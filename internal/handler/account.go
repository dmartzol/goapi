package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/dmartzol/goapi/internal/api"
	"github.com/dmartzol/goapi/internal/proto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createAccount(c *gin.Context) {
	var req api.CreateAccountRequest
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
	c.JSON(http.StatusOK, a.View(vOpts))
}
