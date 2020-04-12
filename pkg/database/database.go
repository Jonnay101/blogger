package database

import (
	"fmt"

	"github.com/Jonnay101/icon/pkg/blog"
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
)

var (
	ErrItemAlreadyExists  = errors.New("item already exists")
	ErrUnknownServerError = errors.New("unknown server error")
)

// Session -
type Session struct {
	*mgo.Session
}

// NewDatabaseSession -
func NewDatabaseSession(mongoURL string) (*Session, error) {

	mgoSession, err := mgo.Dial(mongoURL)
	return &Session{mgoSession}, err
}

// StoreBlogPost - store a blog post in the blog collection
func (s *Session) StoreBlogPost(blogPost *blog.PostData) error {

	fmt.Println(blogPost)
	return nil
}
