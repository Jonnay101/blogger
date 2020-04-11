package blog

import (
	"net/http"
)

// HandlerCreatePost stores a newly created blog post
func (s *server) HandlerCreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var blogPost *Data
		if err := s.decodeRequestBody(w, r, blogPost); err != nil {
			s.respond(w, r, err, http.StatusBadRequest)
			return
		}
	}
}
