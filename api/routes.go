package api

import (
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

var routes = []Route{
	{Path: "/user", Method: http.MethodPost, Handler: HandleCreateUser},
	{Path: "/user", Method: http.MethodGet, Handler: HandleGetAllUser},
	{Path: "/users/{id}", Method: http.MethodDelete, Handler: HandleDeleteUser},
	{Path: "/users/{id}", Method: http.MethodPatch, Handler: HandleUpdateUser},
	{Path: "/patients", Method: http.MethodPost, Handler: HandleCreatePatient},
	{Path: "/patients", Method: http.MethodGet, Handler: HandleGetAllPatient},
	{Path: "/patients/{id}", Method: http.MethodDelete, Handler: HandleDeletePatient},
	{Path: "/patients/{id}", Method: http.MethodPatch, Handler: HandlerUpdatePatient},
	{Path: "/doctors", Method: http.MethodPost, Handler: HandleCreateDoctor},
	{Path: "/doctors", Method: http.MethodGet, Handler: HandleGetAllDoctors},
	{Path: "/doctors/{id}", Method: http.MethodDelete, Handler: HandleDeleteDoctor},
	{Path: "/doctors/{id}", Method: http.MethodPatch, Handler: HandlerUpdateDoctor},
	{Path: "/doctors/{id}/healthinsurence", Method: http.MethodPost, Handler: HandlerAddHealthInsurenceInDoctor},
	{Path: "/doctors/healthinsurence", Method: http.MethodGet, Handler: HandlerGetAllDoctorsWithHealthInsurence},
}
