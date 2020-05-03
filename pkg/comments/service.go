package comments

type database interface {
}

// Service holds all methods for accessing the comments service
type Service struct {
	DB database
}

// NewService -
func NewService(db database) *Service {

	return &Service{db}
}
