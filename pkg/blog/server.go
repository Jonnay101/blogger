package blog

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Server -
type Server interface {
	SetDatabase(db Database)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// Database -
type Database interface {
	StoreBlogPost(*PostData) error
	FindBlogPostByKey(*RequestParams) (*PostData, error)
	FindAllBlogPosts(*RequestParams) ([]*PostData, error)
	UpdateBlogPost(*PostData) error
	RemoveBlogPost(*RequestParams) error
}

type server struct {
	DB     Database
	Router *mux.Router
}

// NewServer -
func NewServer() Server {
	s := &server{}
	s.setRoutes()
	return s
}

func (s *server) SetDatabase(db Database) {
	s.DB = db
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *server) decodeRequestBody(w http.ResponseWriter, r *http.Request, bindObject interface{}) error {
	if r.Body != nil {
		defer r.Body.Close()
	}

	return json.NewDecoder(r.Body).Decode(&bindObject)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, responseData interface{}, statusCode int) {
	w.WriteHeader(statusCode)

	if responseData != nil {
		if err := json.NewEncoder(w).Encode(&responseData); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		}
	}
}
