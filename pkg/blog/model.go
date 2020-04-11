package blog

import (
	"time"

	"github.com/music-tribe/uuid"
)

// Data holds all data and content from a blog postData
type Data struct {
	UUID        uuid.UUID  `json:"_id" bson:"id"`
	Author      string     `json:"author"`
	Title       string     `json:"title"`
	CreatedAt   *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" bson:"updated_at"`
	Content     string     `json:"content"`
	Category    string     `json:"category"`
	StoragePath string     `json:"storage_path"`
	Images      []string   `json:"images"` // a list of image IDs - images already uploaded
}
