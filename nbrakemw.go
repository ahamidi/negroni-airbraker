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

	var n *gobrake.Notice
	res := rw.(negroni.ResponseWriter)
	req := r

	defer func() {
		if res.Status() >= 400 && res.Status() < 600 {
			n = gobrake.NewNotice(errors.New(http.StatusText(res.Status())), req, 20)
			n.Context["environment"] = a.environment
			if req != nil {
				n.Context["uri"] = req.RequestURI
				n.Context["method"] = req.Method
			}
		}
		out, err := a.Notifier.SendNotice(n)
		log.Println("Notice ID:", out)
		if err != nil {
			log.Println("Airbraker Error:", err.Error())
		}
	}()

	next(rw, r)
}
