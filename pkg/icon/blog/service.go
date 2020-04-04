package blog

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
	JSONBuffer, err := storeRouteParamsAsJSONInBuffer(routeParams)
	if err != nil {
		return err
	}
	return e.bindJSONBufferToBlogEntry(JSONBuffer)
}

func storeRouteParamsAsJSONInBuffer(routeParams map[string]string) (*bytes.Buffer, error) {
	JSONBuffer := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(JSONBuffer).Encode(routeParams)
	if err != nil {
		return nil, err
	}
	return JSONBuffer
}

func (e *Entry)bindJSONBufferToBlogEntry(JSONBuffer *bytes.Buffer) error {
	return json.NewDecoder(JSONBuffer).Decode(&e)
}
