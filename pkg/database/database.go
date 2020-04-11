package database

import "github.com/Jonnay101/icon/pkg/blog"

// Session -
type Session struct{}

// NewDatabaseSession -
func NewDatabaseSession() *Session {
	return &Session{}
}

// CreateBlogPost -
func (s *Session) CreateBlogPost(blog.PostData) error {
	// create a new blog post in the database
	return nil
}
