package blog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/music-tribe/uuid"
)

// Entry shows the available methods on the Entry interface
// type Entry interface {
// 	SetEntryFieldsFromRequestBody(r *http.Request) error
// 	ReturnEntryAsJSON() ([]byte, error)
// 	GetStoragePath() string
// }

// Entry holds all data and content from a blog post
type Entry struct {
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

// NewEntry creates a new
func NewEntry() *Entry {
	return &Entry{}
}

// SetEntryFieldsFromRequestBody -
func (e *Entry) SetEntryFieldsFromRequestBody(r *http.Request) error {
	byt, err := ioutil.ReadAll(r.Body)
	if r.Body != nil {
		defer r.Body.Close()
	}
	if err != nil {
		return err
	}
	if err = e.bindJSONDataToBlogEntry(byt); err != nil {
		return err
	}
	e.setUniqueIDForBlogEntry()
	e.setStoragePath()
	return nil
}

func (e *Entry) setStoragePath() {
	e.StoragePath = fmt.Sprintf(
		"/%v/%v/%v/%s",
		e.CreatedAt.Year(),
		e.CreatedAt.Month(),
		e.CreatedAt.Day(),
		e.Title,
	)
}

func (e *Entry) setUniqueIDForBlogEntry() {
	e.UUID = uuid.New()
}

// ReturnEntryAsJSON -
func (e *Entry) ReturnEntryAsJSON() ([]byte, error) {
	return json.Marshal(e)
}

// func encodeRouteParamsToJSONBuffer(routeParams map[string]string) (*bytes.Buffer, error) {
// 	JSONBuffer := bytes.NewBuffer([]byte{})
// 	return JSONBuffer, json.NewEncoder(JSONBuffer).Encode(routeParams)
// }

func (e *Entry) bindJSONDataToBlogEntry(JSONBytes []byte) error {
	return json.NewDecoder(bytes.NewBuffer(JSONBytes)).Decode(&e)
}
