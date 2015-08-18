package negronibrake

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"gopkg.in/airbrake/gobrake.v2"
)

// Middleware is the main airbrake object that sends notifications to the
// Airbrake service
type Middleware struct {
	Notifier    *gobrake.Notifier
	environment string
}

// NewMiddleware returns *Middleware
func NewAirBraker(projectID int64, projectKey string, environment string) *Middleware {
	n := gobrake.NewNotifier(projectID, projectKey)
	return &Middleware{
		Notifier:    n,
		environment: environment,
	}
}

// ServeHTTP is the actual Middleware handler
func (m *Middleware) ServeHTTP(rw http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc) {

	defer func() {
		// Check return code
		res := rw.(negroni.ResponseWriter)
		switch {
		case 400 <= res.Status():
			log.Println("400 Status")
		}
	}()

	next(rw, r)
}
