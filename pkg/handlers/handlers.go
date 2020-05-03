package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jonnay101/icon/pkg/blog"
	"github.com/Jonnay101/icon/pkg/comments"
)

// Handlers -
type Handlers struct {
	Blog     *blog.Service
	Comments *comments.Service
}

// NewHandlers -
func NewHandlers(blogService *blog.Service, commentsService *comments.Service) *Handlers {
	return &Handlers{
		Blog:     blogService,
		Comments: commentsService,
	}
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
