package api

import (
	"net/http"
)

type Route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

var routes = []Route{
	{Path: "/createUniversities", Handler: HandleCreateUniversity},
	{Path: "/getUniversities", Handler: HandleGetAllUniversity},
	{Path: "/deleteUniversities/{id}", Handler: HandleDeleteUniversity},
	{Path: "/updateUniversities/{id}", Handler: HandleUpdateUniversity},
	{Path: "/hello", Handler: HandlerHello},
}
