package common

import (
	"encoding/json"
	"net/http"
)

type CommonController struct{}

type ICommonController interface {
	RespondWithError(w http.ResponseWriter, code int, message string)
	RespondWithJSON(w http.ResponseWriter, code int, payload interface{})
}

func (c *CommonController) RespondWithError(w http.ResponseWriter, code int, message string) {
	c.RespondWithJSON(w, code, map[string]string{"error": message})
}

func (c *CommonController) RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
