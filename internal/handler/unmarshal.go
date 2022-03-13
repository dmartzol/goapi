package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Unmarshal(c *gin.Context, iface interface{}) error {
	b, err := c.GetRawData()
	if err != nil {
		return fmt.Errorf("unable to get raw data: %w", err)
	}
	if h.LogRawRequest {
		h.Infof("payload: %s", b)
	}
	err = json.Unmarshal(b, &iface)
	if err != nil {
		h.Errorf("json.Unmarshal %+v", err)
		return err
	}
	return nil
}
