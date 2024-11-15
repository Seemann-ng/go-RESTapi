package apiserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Seemann-ng/go-RESTapi/internal/app/apiserver"
	"github.com/Seemann-ng/go-RESTapi/internal/app/store/teststore"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := apiserver.NewServer(teststore.New())
	s.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
