package database

import (
	"github.com/Jonnay101/icon/pkg/blog"
	"github.com/globalsign/mgo"
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

	collection := s.DB("omfg").C("blog")
	err := collection.Insert(blogPost)
	if err != nil {
		return err
	}
	return nil
}
