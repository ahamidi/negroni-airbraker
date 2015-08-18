package negronibrake

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {

}

func TestMiddleware(t *testing.T) {
	recorder := httptest.NewRecorder()

	bmw := NewAirBraker(123123, "test", "test")

	n := negroni.New()
	n.Use(bmw)
	n.UseHandler(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(404)
	}))

	n.ServeHTTP(recorder, (*http.Request)(nil))
	assert.Equal(t, recorder.Code, http.StatusNotFound)
}
