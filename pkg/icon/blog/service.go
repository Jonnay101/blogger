package blog

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/music-tribe/uuid"
)

// Entry shows the available methods on the Entry interface
type Entry interface {
	BindRequestParams(r *http.Request) error
}

type entry struct {
	UUID      uuid.UUID  `json:"uuid,omitempty"`
	Author    string     `json:"author,omitempty"`
	Title     string     `json:"title,omitempty"`
	Content   string     `json:"content,omitempty"`
	CreatedAt *time.Time `json:"date,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// BindRequestParams -
func (e *entry) BindRequestParams(r *http.Request) error {
	routeParams := mux.Vars(r)
	JSONBuffer, err := encodeRouteParamsToJSONBuffer(routeParams)
	if err != nil {
		return err
	}
	return e.bindJSONBufferToBlogEntry(JSONBuffer)
}

// BindRequestBody -
func (e *entry) BindRequestBody(r *http.Request) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	defer r.Body.Close()
	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return e.bindJSONBufferToBlogEntry(bytes.NewBuffer(byt))
}

func encodeRouteParamsToJSONBuffer(routeParams map[string]string) (*bytes.Buffer, error) {
	JSONBuffer := bytes.NewBuffer([]byte{})
	return JSONBuffer, json.NewEncoder(JSONBuffer).Encode(routeParams)
}

func (e *entry) bindJSONBufferToBlogEntry(JSONBuffer *bytes.Buffer) error {
	return json.NewDecoder(JSONBuffer).Decode(&e)
}
