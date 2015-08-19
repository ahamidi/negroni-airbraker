package negronibrake

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {

	bmw := NewAirBraker(123123, "test", "test")

	assert.NotNil(t, bmw.Notifier)
	assert.Equal(t, bmw.environment, "test")

}

func TestAirBraker(t *testing.T) {
	recorder := httptest.NewRecorder()

	appID, _ := strconv.Atoi(os.Getenv("AIRBRAKE_APP_ID"))
	bmw := NewAirBraker(int64(appID), os.Getenv("AIRBRAKE_API_KEY"), "development")

	n := negroni.New()
	n.Use(bmw)
	n.UseHandler(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(404)
	}))

	n.ServeHTTP(recorder, (*http.Request)(nil))
	assert.Equal(t, recorder.Code, http.StatusNotFound)

}
