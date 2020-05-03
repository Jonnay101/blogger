package comments

import "github.com/music-tribe/uuid"

// Data holds all the info on a blog post reply
type Data struct {
	BlogPostUUID uuid.UUID `json:"blog_post_uuid"`
	ParentUUID   uuid.UUID `json:"parent_uuid"`
	UserUUID     uuid.UUID `json:"user_uuid"`
	UserName     string    `json:"user_name"`
	Content      string    `json:"content"`
	Likes        string    `json:"likes"`
}

// RequestParams -
type RequestParams struct {
	BlogUUID     uuid.UUID `json:"blog_uuid"`
	BlogPostUUID uuid.UUID `json:"blog_post_uuid"`
	ParentUUID   uuid.UUID `json:"parent_uuid"`
	UserUUID     uuid.UUID `json:"user_uuid"`
	UserName     string    `json:"user_name"`
	Content      string    `json:"content"`
}
