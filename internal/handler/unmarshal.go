package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (h *Handler) Unmarshal(r *http.Request, iface interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Errorf("ReadAll: %+v", err)
		return err
	}
	if h.Verbose {
		h.Infof("payload: %s", body)
	}
	err = json.Unmarshal(body, &iface)
	if err != nil {
		h.Errorf("json.Unmarshal %+v", err)
		return err
	}
	return nil
}
