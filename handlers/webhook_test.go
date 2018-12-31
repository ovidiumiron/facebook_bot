package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebHook(t *testing.T) {
	tt := []struct {
		name      string
		challenge string
		mode      string
		token     string
	}{
		{"first", "foo", "subscribe", "test_token"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			queryString := fmt.Sprintf("hub.verify_token=%s&hub.challenge=%s&hub.mode=%s", tc.token, tc.challenge, tc.mode)
			req, err := http.NewRequest("GET", "localhost:8080/webhook?"+queryString, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()
			WebHook(rec, req)
			fmt.Printf("--->%s\n", string(rec.Body.Bytes()))
		})
	}
}
