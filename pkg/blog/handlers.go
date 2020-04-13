package blog

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Jonnay101/icon/pkg/glitch"
	"github.com/music-tribe/uuid"
)

// HandlerCreatePost stores a newly created blog post
func (s *server) HandlerCreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		blogPost := s.bindRequestBody(w, r)

		s.setBlogPostFields(w, r, blogPost)

		s.storeBlogPost(w, r, blogPost)

		s.respond(w, r, blogPost, http.StatusOK)
		return
	}
}

func (s *server) bindRequestBody(w http.ResponseWriter, r *http.Request) *PostData {

	var blogPost PostData
	if err := s.decodeRequestBody(w, r, &blogPost); err != nil {
		s.respond(w, r, err, http.StatusBadRequest)
		return nil
	}
	return &blogPost
}

func (s *server) setBlogPostFields(w http.ResponseWriter, r *http.Request, blogPost *PostData) {

	blogPost.UUID = uuid.New()

	currentTime := time.Now().UTC().Truncate(time.Second)
	blogPost.CreatedAt = currentTime
	blogPost.UpdatedAt = currentTime
	blogPost.Year = currentTime.Year()
	blogPost.Month = currentTime.Month().String()
	blogPost.Day = currentTime.Day()

	blogPost.DatabaseKey = s.createDatabaseKey(w, r, blogPost)
}

func (s *server) createDatabaseKey(w http.ResponseWriter, r *http.Request, pd *PostData) string {
	if pd.CreatedAt.IsZero() {
		s.respond(w, r, errors.New("PostData CreatedAt field not set"), http.StatusBadRequest)
		return ""
	}

	return fmt.Sprintf(
		"/%d/%s/%d/%s",
		pd.Year,
		pd.Month,
		pd.Day,
		pd.UUID.String(),
	)
}

func (s *server) storeBlogPost(w http.ResponseWriter, r *http.Request, blogPost *PostData) {
	if err := s.DB.StoreBlogPost(blogPost); err != nil {
		if err == glitch.ErrItemAlreadyExists {
			s.respond(w, r, err, http.StatusConflict)
			return
		}
		s.respond(w, r, err, http.StatusInternalServerError)
		return
	}
}
