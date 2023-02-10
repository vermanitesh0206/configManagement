package controllers

import (
	"bytes"
	"encoding/json"
	. "my-app/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {

	validCred := Credentials{
		Username: "user1",
		Password: "password1",
	}
	invalidCred := Credentials{
		Username: "dummy",
		Password: "dummy",
	}
	var testCases = []struct {
		scenario     string
		cred         Credentials
		expectedCode int
	}{
		{"Valid User", validCred, 200},
		{"Invalid User", invalidCred, 401},
	}

	for _, tt := range testCases {
		body, _ := json.Marshal(tt.cred)
		req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Login)
		handler.ServeHTTP(rr, req)

		actualCode := rr.Code
		if actualCode != tt.expectedCode {
			t.Errorf("Scenario(%s): expected %d, actual %d", tt.scenario, tt.expectedCode, actualCode)
		}
	}
}
