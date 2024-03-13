package api

import (
	"net/http"
)

type Route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

var routes = []Route{
	{Path: "/user", Handler: HandleCreateUser},
	{Path: "/user", Handler: HandleGetAllUser},
	{Path: "/users/{id}", Handler: HandleDeleteUser},
	{Path: "/users/{id}", Handler: HandleUpdateUser},
	{Path: "/doctors", Handler: HandleCreateDoctor},
	{Path: "/doctors", Handler: HandleGetAllDoctors},
	{Path: "/doctors/{id}", Handler: HandleDeleteDoctor},
	{Path: "/doctors/{id}", Handler: HandlerUpdateDoctor},
	{Path: "/doctors/healthinsurence/{id}", Handler: HandlerAddHealthInsurenceInDoctor},
	{Path: "/doctors/healthinsurence", Handler: HandlerGetAllDoctorsWithHealthInsurence},
}
