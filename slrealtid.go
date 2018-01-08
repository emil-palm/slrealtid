package slrealtid

import (
	"net/http"
	"github.com/mrevilme/slrealtid/versions/base"
)

type client interface {
	Departures() []base.Departures
	DeparturesNow() []base.Departures
}

type Client struct {
	client
	Client http.Client
	Key string
}