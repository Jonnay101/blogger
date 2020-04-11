package blog

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type database interface {
	CreateBlogPost(PostData) error
}

type server struct {
	DB     database
	Router *mux.Router
}

// NewServer -
func NewServer() http.Handler {
	s := &server{}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, responseData interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	if responseData != nil {
		if err := json.NewEncoder(w).Encode(iface); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		}
	}
}

func (s *server) decodeRequestBody(w http.ResponseWriter, r *http.Request, iface interface{}) error {
	if r.Body != nil {
		defer r.Body.Close()
	}
	return json.NewDecoder(r.Body).Decode(&iface)
}
