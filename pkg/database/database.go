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
func (s *Session) FindBlogPostByID(id string) (*blog.PostData, error) {

	blogPosts := s.getBlogCollection()
	blogPost := &blog.PostData{}

	q := blogPosts.FindId(id)
	if err := q.One(&blogPost); err != nil {
		if err == mgo.ErrNotFound {
			return blogPost, glitch.ErrRecordNotFound
		}
		return blogPost, err
	}

	return blogPost, nil
}

// FindAllBlogPosts - finds all posts that match the fields in the query map
func (s *Session) FindAllBlogPosts(reqParams *blog.RequestParams) ([]*blog.PostData, error) {

	blogPosts := s.getBlogCollection()
	matchingBlogPosts := []*blog.PostData{}

	q := blogPosts.Find(reqParams.QueryMap)
	if err := q.All(&matchingBlogPosts); err != nil {
		if err == mgo.ErrNotFound {
			return matchingBlogPosts, glitch.ErrRecordNotFound
		}
		return matchingBlogPosts, err
	}

	return matchingBlogPosts, nil
}

// UpdateBlogPost - updates the post with the corresponding id
func (s *Session) UpdateBlogPost(blogPost *blog.PostData) error {

	blogPosts := s.getBlogCollection()

	if err := blogPosts.UpdateId(blogPost.DatabaseKey, blogPost); err != nil {
		if err == mgo.ErrNotFound {
			return glitch.ErrRecordNotFound
		}
		return err
	}

	return nil
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
