package database

import (
	"github.com/Jonnay101/icon/pkg/blog"
	"github.com/Jonnay101/icon/pkg/glitch"
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

func (s *Session) getBlogCollection() *mgo.Collection {

	return s.DB("omfg").C("blog")
}

// StoreBlogPost - store a blog post in the blog collection
func (s *Session) StoreBlogPost(blogPost *blog.PostData) error {

	blogPosts := s.getBlogCollection()

	q := blogPosts.FindId(blogPost.DatabaseKey)
	if n, _ := q.Count(); n > 0 {
		return glitch.ErrItemAlreadyExists
	}

	return blogPosts.Insert(blogPost)
}

// FindBlogPostByID - find a single blog post using the id
func (s *Session) FindBlogPostByID(reqParams *blog.RequestParams) (*blog.PostData, error) {

	blogPosts := s.getBlogCollection()
	blogPost := &blog.PostData{}

	q := blogPosts.FindId(reqParams.DatabaseKey)
	if err := q.One(&blogPost); err != nil {
		if err == mgo.ErrNotFound {
			return nil, glitch.ErrRecordNotFound
		}
		return nil, err
	}

	return blogPost, nil
}

// FindAllBlogPosts - finds all posts that match the fields in the query map
func (s *Session) FindAllBlogPosts(reqParams *blog.RequestParams) ([]*blog.PostData, error) {

	blogPosts := s.getBlogCollection()
	var matchingBlogPosts []*blog.PostData

	q := blogPosts.Find(reqParams.QueryMap)
	if err := q.All(&matchingBlogPosts); err != nil {
		if err == mgo.ErrNotFound {
			return nil, glitch.ErrRecordNotFound
		}
		return nil, err
	}

	return matchingBlogPosts, nil
}

// RemoveBlogPost - removes the post matching the id
func (s *Session) RemoveBlogPost(reqParams *blog.RequestParams) error {

	blogPosts := s.getBlogCollection()

	if err := blogPosts.RemoveId(reqParams.DatabaseKey); err != nil {
		if err == mgo.ErrNotFound {
			return glitch.ErrRecordNotFound
		}
		return err
	}

	return nil
}
