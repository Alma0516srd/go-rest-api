package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiServer_HandleHello(t *testing.T) {
	apiserver := New(NewConfig())
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	apiserver.handleHello().ServeHTTP(recorder, request)
	assert.Equal(t, "Hello World", recorder.Body.String())
}
