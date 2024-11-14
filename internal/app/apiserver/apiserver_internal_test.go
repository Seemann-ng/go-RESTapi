package apiserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Seemann-ng/go-RESTapi/internal/app/apiserver"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := apiserver.New(apiserver.NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	s.HandleHello().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello World", rec.Body.String())
}
