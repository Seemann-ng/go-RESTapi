package apiserver_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"

	"github.com/Seemann-ng/go-RESTapi/internal/app/apiserver"
	"github.com/Seemann-ng/go-RESTapi/internal/app/model"
	"github.com/Seemann-ng/go-RESTapi/internal/app/store/teststore"
)

func TestServer_HandleUsersCreate(test *testing.T) {
	server := apiserver.NewServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "test@example.com",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": "password",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    "test@example.com",
				"password": "inv",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no email",
			payload: map[string]string{
				"password": "password",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "no password",
			payload: map[string]string{
				"email": "test@example.com",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			recorder := httptest.NewRecorder()
			buffer := &bytes.Buffer{}
			if err := json.NewEncoder(buffer).Encode(testCase.payload); err != nil {
				test.Fatal(err)
			}

			request, _ := http.NewRequest(http.MethodPost, "/users", buffer)
			server.ServeHTTP(recorder, request)

			assert.Equal(test, testCase.expectedCode, recorder.Code)
		})
	}
}

func TestServer_handleSessionCreate(test *testing.T) {
	testUser := model.TestUser(test)
	storage := teststore.New()
	if err := storage.User().Create(testUser); err != nil {
		test.Fatal(err)
	}

	server := apiserver.NewServer(storage, sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    testUser.Email,
				"password": testUser.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": testUser.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    testUser.Email,
				"password": "inv",
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "no email",
			payload: map[string]string{
				"password": testUser.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "no password",
			payload: map[string]string{
				"email": testUser.Email,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "no data",
			payload:      nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			recorder := httptest.NewRecorder()
			buffer := &bytes.Buffer{}
			if err := json.NewEncoder(buffer).Encode(testCase.payload); err != nil {
				test.Fatal(err)
			}

			request, _ := http.NewRequest(http.MethodPost, "/sessions", buffer)
			server.ServeHTTP(recorder, request)

			assert.Equal(test, testCase.expectedCode, recorder.Code)
		})
	}
}
