package help

import (
	"encoding/json"
	"net/http"
	"time"
)

// GetCurrentUTCTime returns the current Universal Time to the nearest second
func GetCurrentUTCTime() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}

// DecodeRequestBody decodes the json request body into the provided binding object
func DecodeRequestBody(w http.ResponseWriter, r *http.Request, bindObject interface{}) error {

	if r.Body != nil {
		defer r.Body.Close()
	}

	return json.NewDecoder(r.Body).Decode(&bindObject)
}
