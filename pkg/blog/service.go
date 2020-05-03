package blog

type database interface {
	StoreBlogPost(*PostData) error
	FindBlogPostByKey(*RequestParams) (*PostData, error)
	FindAllBlogPosts(*RequestParams) ([]*PostData, error)
	UpdateBlogPost(*PostData) error
	RemoveBlogPost(*RequestParams) error
}

type Service struct {
	DB database
}

// NewService -
func NewService(db database) *Service {
	return &Service{db}
}
