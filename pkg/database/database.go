package database

import (
	"fmt"

	"github.com/Jonnay101/icon/pkg/blog"
	"github.com/pkg/errors"
)

var (
	ErrItemAlreadyExists  = errors.New("item already exists")
	ErrUnknownServerError = errors.New("unknown server error")
)

// Session -
type Session struct{}

// NewDatabaseSession -
func NewDatabaseSession() (*Session, error) {
	return &Session{}, nil
}

// StoreBlogPost -
func (s *Session) StoreBlogPost(blogPost *blog.PostData) error {
	// create a new blog post in the database
	fmt.Println(blogPost)
	return nil
}
