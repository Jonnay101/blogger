package blog

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Jonnay101/icon/pkg/glitch"
	"github.com/gorilla/mux"
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

func (s *server) HandlerGetPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		requestParams := s.bindRequestParams(w, r)
		blogPost := s.findBlogPost(w, r, requestParams)
		s.respond(w, r, blogPost, http.StatusOK)
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

func (s *server) bindRequestParams(w http.ResponseWriter, r *http.Request) *RequestParams {

	var requestParams RequestParams
	dynamicRoutes := mux.Vars(r)
	requestParams.DatabaseKey = strings.TrimPrefix(r.URL.Path, "/blog")
	requestParams.Category = r.URL.Query().Get("category")
	requestParams.Year = dynamicRoutes["year"]
	requestParams.Month = dynamicRoutes["month"]
	requestParams.Day = dynamicRoutes["day"]

	return &requestParams
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

func (s *server) findBlogPost(w http.ResponseWriter, r *http.Request, reqParams *RequestParams) *PostData {

	blogPost, err := s.DB.FindBlogPostByID(reqParams)
	if err != nil {
		if err == glitch.ErrRecordNotFound {
			s.respond(w, r, err, http.StatusNotFound)
			return nil
		}
		s.respond(w, r, err, http.StatusInternalServerError)
		return nil
	}

	return blogPost
}
