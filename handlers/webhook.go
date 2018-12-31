package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

type HubModeError struct{}

func (e HubModeError) Error() string {
	return "HubModeError"
}

type ChallengeError struct{}

func (e ChallengeError) Error() string {
	return "ChallengeError"
}

type VerifyTokenError struct{}

func (e VerifyTokenError) Error() string {
	return "VerifyTokenError"
}
func verifyToken(v url.Values) (bool, error) {
	const token = "test_token"
	t, ok := v["hub.verify_token"]
	if !ok || len(t[0]) < 1 {
		return false, VerifyTokenError{}
	}
	return t[0] == token, nil
}
func getHubMode(v url.Values) (string, error) {
	hubMode, ok := v["hub.mode"]

	if !ok || len(hubMode[0]) < 1 {
		return "", HubModeError{}
	}
	return hubMode[0], nil
}
func getChallenge(v url.Values) (string, error) {
	challenge, ok := v["hub.challenge"]
	if !ok || len(challenge[0]) < 1 {
		return "", ChallengeError{}
	}
	return challenge[0], nil
}

// WebHook is the handler
func WebHook(w http.ResponseWriter, r *http.Request) {
	if r != nil && r.Body != nil {
		defer r.Body.Close()
	}
	values := r.URL.Query()
	isAllowed, err := verifyToken(values)
	if err != nil {
		// TODO: log
		return
	}
	if !isAllowed {
		return
	}
	hubMode, err := getHubMode(values)
	if err != nil {
		return
	}
	if hubMode == "subscribe" {
		c, err := getChallenge(values)
		if err != nil {
			return
		}
		fmt.Fprintf(w, "%s", c)
	}
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()

	fmt.Printf("this is the body: -->%s<---s\n", newStr)
}
