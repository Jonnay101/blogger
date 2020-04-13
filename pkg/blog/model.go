package blog

import (
	"time"

	"github.com/music-tribe/uuid"
)

// PostData holds all data and content from a blog postData
type PostData struct {
	DatabaseKey string    `json:"_id" bson:"_id"`
	UUID        uuid.UUID `json:"uuid" bson:"uuid"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	Year        int       `json:"year"`
	Month       string    `json:"month"`
	Day         int       `json:"day"`
	Category    string    `json:"category"`
	Metadata    []string  `json:"metadata"`
	Images      []string  `json:"images"` // a list of image IDs - images already uploaded
}
