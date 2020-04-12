package database

import (
	"fmt"

	"github.com/Jonnay101/icon/pkg/blog"
)

// Session -
type Session struct{}

// NewDatabaseSession -
func NewDatabaseSession() (*Session, error) {
	return &Session{}, nil
}

// CreateBlogPost -
func (s *Session) CreateBlogPost(blogPost *blog.PostData) error {
	// create a new blog post in the database
	fmt.Println(blogPost)
	return nil
}
