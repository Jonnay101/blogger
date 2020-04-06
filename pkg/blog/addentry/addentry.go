package addentry

import (
	"net/http"

	"github.com/Jonnay101/icon/pkg/blog"
	"github.com/Jonnay101/icon/pkg/glitch"
)

// methods required for this handler
type database interface {
	AddBlogEntry(*blog.Entry) error
}

// handler satisfies http.Handler interface with ServeHTTP
type handler struct {
	db database
}

// NewHandler -
func NewHandler(db database) http.Handler {
	return &handler{db}
}

// ServeHTTP - required to make Handler an http.Handler
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	blogEntry := blog.NewEntry()

	if err := blogEntry.SetEntryFieldsFromRequestBody(r); err != nil {
		glitch.LogError(err)
	}

	if err := h.db.AddBlogEntry(blogEntry); err != nil {
		glitch.LogError(err)
	}

	entryJSON, err := blogEntry.ReturnEntryAsJSON()
	if err != nil {
		glitch.LogError(err)
	}

	w.Write(entryJSON)
}
