package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golangApi/https-rest-api/internal/app/model"
	"golangApi/https-rest-api/internal/app/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.New()
	u := model.TestUser(t)
	store.User().Create(u)

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	secretKey := []byte("secret")
	s := NewServer(store, sessions.NewCookieStore(secretKey))
	sc := securecookie.New([]byte("secret"), nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			cookieStr, _ := sc.Encode(sessionStore, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s = %s", sessionStore, cookieStr))
			s.authenticalUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServerHandelUsersCreate(t *testing.T) {
	s := NewServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		playload     interface{}
		expectedCode int
	}{
		{
			name: "valid",
			playload: map[string]string{
				"email":    "user@example",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid playload",
			playload:     "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			playload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.playload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := NewServer(store)
	testCases := []struct {
		name         string
		playload     interface{}
		expectedCode int
	}{
		{
			name: "valid",
			playload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid playlod",
			playload:     "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			playload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			playload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.playload)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
