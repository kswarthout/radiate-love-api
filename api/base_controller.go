package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// BaseController handle all base methods
type BaseController struct {
}

// Error is a struct to return an auth error
type Error struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// SendJSON marshals v to a json struct and send the appropriate headers to w
func (c *BaseController) SendJSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")

	b, err := json.Marshal(v)

	if err != nil {
		log.Print(fmt.Sprintf("Error while encoding JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error"}: "Internal server error"`)
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(b))
	}
}
