package blog

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Jonnay101/icon/pkg/glitch"
	"github.com/gorilla/mux"
	"github.com/music-tribe/uuid"
	"gopkg.in/mgo.v2/bson"
)

// HandlerCreatePost stores a newly created blog post
func (s *server) HandlerCreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		blogPost, err := s.bindRequestBody(w, r)
		if err != nil {
			s.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		if err = s.validateBlogPostRequest(blogPost); err != nil {
			s.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		if err = s.initBlogPostData(w, r, blogPost); err != nil {
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := s.DB.StoreBlogPost(blogPost); err != nil {
			if err == glitch.ErrItemAlreadyExists {
				s.respond(w, r, err.Error(), http.StatusConflict)
				return
			}
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		s.respond(w, r, blogPost, http.StatusOK)
	}
}

func (s *server) HandlerGetPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqParams, err := s.bindRequestParams(w, r)
		if err != nil {
			s.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		blogPost, err := s.DB.FindBlogPostByKey(reqParams)
		if err != nil {
			if err == glitch.ErrRecordNotFound {
				s.respond(w, r, err.Error(), http.StatusNotFound)
				return
			}
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		s.respond(w, r, blogPost, http.StatusOK)
	}
}

func (s *server) HandlerGetAllPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqParams, err := s.bindRequestParams(w, r)
		if err != nil {
			s.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		blogPosts, err := s.DB.FindAllBlogPosts(reqParams)
		if err != nil {
			if err == glitch.ErrRecordNotFound {
				s.respond(w, r, err.Error(), http.StatusNotFound)
				return
			}
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		s.respond(w, r, blogPosts, http.StatusOK)
	}
}

func (s *server) HandlerUpdatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqParams, err := s.bindRequestParams(w, r)
		if err != nil {
			s.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		oldBlogPost, err := s.DB.FindBlogPostByKey(reqParams)
		if err != nil {
			if err == glitch.ErrRecordNotFound {
				s.respond(w, r, err.Error(), http.StatusNotFound)
				return
			}
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
		}

		newBlogPost, err := s.bindRequestBody(w, r)
		if err != nil {
			s.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		populateZeroValueFieldsWithOldData(oldBlogPost, newBlogPost)

		if err := s.DB.UpdateBlogPost(newBlogPost); err != nil {
			if err == glitch.ErrRecordNotFound {
				s.respond(w, r, err.Error(), http.StatusNotFound)
				return
			}
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		s.respond(w, r, newBlogPost, http.StatusOK)
	}
}

func (s *server) HandlerDeletePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqParams, err := s.bindRequestParams(w, r)
		if err != nil {
			s.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.DB.RemoveBlogPost(reqParams); err != nil {
			if err == glitch.ErrRecordNotFound {
				s.respond(w, r, err.Error(), http.StatusNotFound)
				return
			}
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		s.respond(w, r, "OK", http.StatusOK)
	}
}

func (s *server) bindRequestBody(w http.ResponseWriter, r *http.Request) (*PostData, error) {

	var blogPost PostData
	if err := s.decodeRequestBody(w, r, &blogPost); err != nil {
		return nil, err
	}

	dynamicRoutes := mux.Vars(r)
	var err error

	blogPost.UserUUID, err = uuid.Parse(dynamicRoutes["user_uuid"])
	if err != nil {
		return nil, err
	}

	if blogUUID, ok := dynamicRoutes["uuid"]; ok {
		blogPost.UUID, err = uuid.Parse(blogUUID)
		if err != nil {
			return nil, err
		}

		blogPost.DatabaseKey = getDatabaseKeyFromURLPath(r, blogPost.UserUUID)
	}

	return &blogPost, nil
}

func (s *server) bindRequestParams(w http.ResponseWriter, r *http.Request) (*RequestParams, error) {

	var reqParams RequestParams
	var err error
	dynamicRoutes := mux.Vars(r)

	if uuidStr, ok := dynamicRoutes["uuid"]; ok {
		reqParams.UUID, err = uuid.Parse(uuidStr)
		if err != nil {
			return nil, err
		}
	}

	reqParams.UserUUID, err = uuid.Parse(dynamicRoutes["user_uuid"])
	if err != nil {
		return nil, err
	}

	reqParams.DatabaseKey = getDatabaseKeyFromURLPath(r, reqParams.UserUUID)

	reqParams.Year, err = getRequestParamInt(w, r, "year")
	if err != nil {
		return nil, err
	}

	reqParams.Month, err = getRequestParamString(w, r, "month")
	if err != nil {
		return nil, err
	}

	reqParams.Day, err = getRequestParamInt(w, r, "day")
	if err != nil {
		return nil, err
	}

	queries := r.URL.Query()
	reqParams.Author = queries.Get("author")
	reqParams.Title = queries.Get("title")
	reqParams.Category = queries.Get("category")

	if err := setQueryConfig(&reqParams); err != nil {
		return nil, err
	}

	return &reqParams, nil
}

func setQueryConfig(reqParams *RequestParams) error {

	reqParams.QueryConfig = bson.M{}

	if reqParams.UUID != uuid.Nil {
		// when the uuid is passed in the url, only the _id needs to be queried
		reqParams.QueryConfig["_id"] = reqParams.DatabaseKey
		return nil
	}

	if reqParams.Year != 0 {
		reqParams.QueryConfig["year"] = reqParams.Year
	}

	if reqParams.Month != "" {
		reqParams.QueryConfig["month"] = reqParams.Month
	}

	if reqParams.Day != 0 {
		reqParams.QueryConfig["day"] = reqParams.Day
	}

	if reqParams.Title != "" {
		reqParams.QueryConfig["title"] = reqParams.Title
	}

	if reqParams.Author != "" {
		reqParams.QueryConfig["author"] = reqParams.Author
	}

	if reqParams.Category != "" {
		reqParams.QueryConfig["category"] = reqParams.Category
	}

	return nil
}

func (s *server) initBlogPostData(w http.ResponseWriter, r *http.Request, blogPost *PostData) error {

	blogPost.UUID = uuid.New()

	s.setDateCreatedAt(blogPost)

	err := s.createDatabaseKey(blogPost)

	return err
}

func (*server) setDateCreatedAt(blogPost *PostData) {

	currentTime := getCurrentUTCTime()

	blogPost.CreatedAt = currentTime
	blogPost.UpdatedAt = currentTime
	blogPost.Year = currentTime.Year()
	blogPost.Month = currentTime.Month().String()
	blogPost.Day = currentTime.Day()
}

func (*server) createDatabaseKey(blogPost *PostData) error {

	if blogPost.CreatedAt.IsZero() {
		return errors.New("PostData CreatedAt field not set")
	}

	blogPost.DatabaseKey = fmt.Sprintf(
		"/%d/%s/%d/%s",
		blogPost.Year,
		blogPost.Month,
		blogPost.Day,
		blogPost.UUID.String(),
	)

	return nil
}

func getRequestParamInt(w http.ResponseWriter, r *http.Request, param string) (int, error) {

	var num int
	var err error
	routeParams := mux.Vars(r)
	queries := r.URL.Query()

	if item, ok := routeParams[param]; ok {
		num, err = strconv.Atoi(item)
		if err != nil {
			return 0, err
		}
	}

	if item, ok := queries[param]; ok {
		num, err = strconv.Atoi(item[0])
		if err != nil {
			return 0, err
		}
	}

	return num, nil
}

func getRequestParamString(w http.ResponseWriter, r *http.Request, param string) (string, error) {

	var str string
	routeParams := mux.Vars(r)
	queries := r.URL.Query()

	if item, ok := routeParams[param]; ok {
		str = item
	}

	if item, ok := queries[param]; ok {
		str = item[0]
	}

	return str, nil
}

func getDatabaseKeyFromURLPath(r *http.Request, blogUserUUID uuid.UUID) string {
	return strings.TrimPrefix(r.URL.Path, fmt.Sprintf("/blog/%s", blogUserUUID.String()))
}

func (s *server) validateBlogPostRequest(blogPost *PostData) error {

	if (blogPost.Author) == "" {
		return errors.New("blog post 'author' field must not be empty")
	}

	if (blogPost.Title) == "" {
		return errors.New("blog post 'title' field must not be empty")
	}

	if blogPost.Category == "" {
		return errors.New("blog post 'category' field must not be empty")
	}

	if blogPost.Content == "" {
		return errors.New("blog post 'content' field must not be empty")
	}

	return nil
}

func populateZeroValueFieldsWithOldData(oldBlogPost, newBlogPost *PostData) {

	// compare the 2 objects and decipher which fields need replacing
	if newBlogPost.UUID == uuid.Nil {
		newBlogPost.UUID = oldBlogPost.UUID
	}

	if newBlogPost.Author == "" {
		newBlogPost.Author = oldBlogPost.Author
	}

	if newBlogPost.Title == "" {
		newBlogPost.Title = oldBlogPost.Title
	}

	if newBlogPost.Content == "" {
		newBlogPost.Content = oldBlogPost.Content
	}

	if newBlogPost.Category == "" {
		newBlogPost.Category = oldBlogPost.Category
	}

	if newBlogPost.Metadata == nil {
		newBlogPost.Metadata = oldBlogPost.Metadata
	}

	if newBlogPost.Images == nil {
		newBlogPost.Images = oldBlogPost.Images
	}

	newBlogPost.CreatedAt = oldBlogPost.CreatedAt
	newBlogPost.Year = oldBlogPost.Year
	newBlogPost.Month = oldBlogPost.Month
	newBlogPost.Day = oldBlogPost.Day
	newBlogPost.UpdatedAt = getCurrentUTCTime()
}
