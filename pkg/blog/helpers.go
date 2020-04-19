package blog

import (
	"time"

	"github.com/music-tribe/uuid"
)

func getCurrentUTCTime() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}

func replaceZeroValueFieldsWithOldData(oldBlogPost, newBlogPost *PostData) {

	// compare the 2 objects and decipher which fields need replacing
	if newBlogPost.UUID == uuid.Nil {
		newBlogPost.UUID = oldBlogPost.UUID
	}

	if newBlogPost.Author == "" {
		newBlogPost.Author = oldBlogPost.Author
	}

	if newBlogPost.Title == "" {
		newBlogPost.Title = oldBlogPost.Title
	}

	if newBlogPost.Content == "" {
		newBlogPost.Content = oldBlogPost.Content
	}

	if newBlogPost.Category == "" {
		newBlogPost.Category = oldBlogPost.Category
	}

	if newBlogPost.Metadata == nil {
		newBlogPost.Metadata = oldBlogPost.Metadata
	}

	if newBlogPost.Images == nil {
		newBlogPost.Images = oldBlogPost.Images
	}

	newBlogPost.CreatedAt = oldBlogPost.CreatedAt
	newBlogPost.Year = oldBlogPost.Year
	newBlogPost.Month = oldBlogPost.Month
	newBlogPost.Day = oldBlogPost.Day
	newBlogPost.UpdatedAt = getCurrentUTCTime()
}
