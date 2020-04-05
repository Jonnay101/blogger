package blog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/music-tribe/uuid"
)

// Entry shows the available methods on the Entry interface
type Entry interface {
	SetEntryFieldsFromRequestBody(r *http.Request) error
	SetBlogEntryID(uuid.UUID)
	ReturnEntryAsJSON() ([]byte, error)
}

type entry struct {
	UUID      uuid.UUID  `json:"uuid,omitempty"`
	Author    string     `json:"author,omitempty"`
	Title     string     `json:"title,omitempty"`
	Content   string     `json:"content,omitempty"`
	CreatedAt *time.Time `json:"date,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewEntry creates a new
func NewEntry() Entry {
	return &entry{}
}

func (e *entry) SetEntryFieldsFromRequestBody(r *http.Request) error {
	byt, err := ioutil.ReadAll(r.Body)
	if r.Body != nil {
		defer r.Body.Close()
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(byt, &e)
}

func (e *entry) SetBlogEntryID(id uuid.UUID) {
	e.UUID = id
}

func (e *entry) ReturnEntryAsJSON() ([]byte, error) {
	return json.Marshal(e)
}

// // BindRouteParamsQueriesAndBody
// func (e *entry) BindRouteParamsQueriesAndBody(r *http.Request) error {
// 	if err := e.BindRequestParams(r); err != nil {
// 		return err
// 	}
// 	if err := e.BindRequestQueries(r); err != nil {
// 		return err
// 	}
// 	return e.BindRequestBody(r)
// }

// // BindRequestParams -
// func (e *entry) BindRequestParams(r *http.Request) error {
// 	routeParams := mux.Vars(r)
// 	JSONBuffer, err := encodeRouteParamsToJSONBuffer(routeParams)
// 	if err != nil {
// 		return err
// 	}
// 	return e.bindJSONBufferToBlogEntry(JSONBuffer)
// }

// // BindRequestBody -
// func (e *entry) BindRequestBody(r *http.Request) error {
// 	if r.Body == nil {
// 		return errors.New("request body is empty")
// 	}
// 	defer r.Body.Close()
// 	byt, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return err
// 	}
// 	return e.bindJSONBufferToBlogEntry(bytes.NewBuffer(byt))
// }

// // BindRequestQueries -
// func (e *entry) BindRequestQueries(r *http.Request) error {
// 	queriesReadyForBinding := flattenQueryMapValues(r.URL.Query())
// 	JSONBuffer, err := encodeRouteParamsToJSONBuffer(queriesReadyForBinding)
// 	if err != nil {
// 		return err
// 	}
// 	return e.bindJSONBufferToBlogEntry(JSONBuffer)
// }

// func flattenQueryMapValues(queryMap map[string][]string) map[string]string {
// 	flatQueryMap := make(map[string]string)
// 	for key, value := range queryMap {
// 		flatQueryMap[key] = value[0]
// 	}
// 	return flatQueryMap
// }

// func encodeRouteParamsToJSONBuffer(routeParams map[string]string) (*bytes.Buffer, error) {
// 	JSONBuffer := bytes.NewBuffer([]byte{})
// 	return JSONBuffer, json.NewEncoder(JSONBuffer).Encode(routeParams)
// }

// func (e *entry) bindJSONBufferToBlogEntry(JSONBuffer *bytes.Buffer) error {
// 	return json.NewDecoder(JSONBuffer).Decode(&e)
// }
