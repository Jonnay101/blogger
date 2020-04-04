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

// Entry -
type Entry struct {
	UUID      uuid.UUID  `json:"uuid,omitempty"`
	Author    string     `json:"author,omitempty"`
	Title     string     `json:"title,omitempty"`
	Content   string     `json:"content,omitempty"`
	CreatedAt *time.Time `json:"date,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// BindRequestParams -
func (e *Entry) BindRequestParams(r *http.Request) error {
	routeParams := mux.Vars(r)
	JSONBuffer, err := encodeRouteParamsToJSONBuffer(routeParams)
	if err != nil {
		return err
	}
	return e.bindJSONBufferToBlogEntry(JSONBuffer)
}

// BindRequestBody -
func (e *Entry) BindRequestBody(r *http.Request) error {
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
	err := json.NewEncoder(JSONBuffer).Encode(routeParams)
	if err != nil {
		return nil, err
	}
	return JSONBuffer, nil
}

func (e *Entry) bindJSONBufferToBlogEntry(JSONBuffer *bytes.Buffer) error {
	return json.NewDecoder(JSONBuffer).Decode(&e)
}
