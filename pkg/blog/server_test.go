package blog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/music-tribe/uuid"
)

var userUUID uuid.UUID = uuid.New()

var body = map[string]string{
	"user_uuid": userUUID.String(),
	"author":    "testName",
	"title":     "testTitle",
}

func Test_server_decodeRequestBody(t *testing.T) {

	byt, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	mockRT := &mux.Router{}
	mockDB := &MockDatabase{}
	body := bytes.NewReader(byt)

	type fields struct {
		DB     Database
		Router *mux.Router
	}

	type args struct {
		w          http.ResponseWriter
		r          *http.Request
		bindObject interface{}
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"invalid user uuid",
			fields{
				mockDB,
				mockRT,
			},
			args{
				httptest.NewRecorder(),
				httptest.NewRequest("get", "/", body),
				&PostData{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				DB:     tt.fields.DB,
				Router: tt.fields.Router,
			}

			if err := s.decodeRequestBody(tt.args.w, tt.args.r, tt.args.bindObject); (err != nil) != tt.wantErr {
				t.Errorf("server.decodeRequestBody() error = %v, wantErr %v", err, tt.wantErr)
			}

			if blogPost, ok := tt.args.bindObject.(*PostData); ok && blogPost.UserUUID != userUUID {
				t.Errorf("Expected UserUUID to be %s, got %s", userUUID, blogPost.UserUUID)
			}

			fmt.Printf("%+v", tt.args.bindObject)
		})
	}
}
