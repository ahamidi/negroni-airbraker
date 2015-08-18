package negronibrake

import (
	"errors"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"gopkg.in/airbrake/gobrake.v2"
)

// AirBraker is the main airbrake object that sends notifications to the
// Airbrake service
type AirBraker struct {
	Notifier    *gobrake.Notifier
	environment string
}

// NewAirBraker returns *AirBraker
func NewAirBraker(projectID int64, projectKey string, environment string) *AirBraker {
	n := gobrake.NewNotifier(projectID, projectKey)
	return &AirBraker{
		Notifier:    n,
		environment: environment,
	}
}

// ServeHTTP is the actual AirBraker handler
func (a *AirBraker) ServeHTTP(rw http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc) {

	// Check return code
	res := rw.(negroni.ResponseWriter)
	req := r

	var n *gobrake.Notice
	if res.Status() >= 400 && res.Status() < 600 {
		n := gobrake.NewNotice(errors.New(http.StatusText(res.Status())), r, 1)
		n.Context["environment"] = a.environment
		if req != nil {
			n.Context["uri"] = req.RequestURI
			n.Context["method"] = req.Method
		}
		_, err := a.Notifier.SendNotice(n)
		if err != nil {
			log.Println("Airbraker Error:", err.Error())
		}
	}

	next(rw, r)

	defer a.Notifier.Flush()
	_, err := a.Notifier.SendNotice(n)
	if err != nil {
		log.Println("Airbraker Error:", err.Error())
	}

}
