package add-blog-entry

// methods required for this handler
type database interface {
	AddBlogEntry(blog.Entry) error
}

// Handler satisfies http.Handler interface with ServeHTTP
type Handler struct {
	db database
}

// NewHandler - 
func NewHandler(db database) *http.Handler {
	return &Handler{db}
}

// ServeHTTP - required to make Handler an http.Handler
func (h *Handler)ServeHTTP(w http.ResponseWriter,r *http.Request) {
	// bind request to blog.Entry struct
	var blogEntry = 
	h.db.AddBlogEntry()
}

func bindRequestParamsToBlogEntry()