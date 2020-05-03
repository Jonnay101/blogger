package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jonnay101/icon/pkg/blog"
)

// Handlers -
type Handlers struct {
	Blog *blog.Service
}

// NewHandlers -
func NewHandlers(blogService *blog.Service) *Handlers {
	return &Handlers{
		blogService,
	}
}

func (h *Handlers) decodeRequestBody(w http.ResponseWriter, r *http.Request, bindObject interface{}) error {

	if r.Body != nil {
		defer r.Body.Close()
	}

	return json.NewDecoder(r.Body).Decode(&bindObject)
}

func (h *Handlers) respond(w http.ResponseWriter, r *http.Request, responseData interface{}, statusCode int) {

	w.WriteHeader(statusCode)

	if responseData != nil {
		if err := json.NewEncoder(w).Encode(&responseData); err != nil {
			// TODO: check this - seems wrong
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		}
	}
}
