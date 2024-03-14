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
	{Path: "/admin", Method: http.MethodPost, Handler: HandleCreateAdmin},
	{Path: "/admin", Method: http.MethodGet, Handler: HandleGetAllAdmins},
	{Path: "/admins/{id}", Method: http.MethodDelete, Handler: HandleDeleteAdmin},
	{Path: "/admins/{id}", Method: http.MethodPatch, Handler: HandleUpdateAdmin},
	{Path: "/patients", Method: http.MethodPost, Handler: HandleCreatePatient},
	{Path: "/patients", Method: http.MethodGet, Handler: HandleGetAllPatient},
	{Path: "/patients/{id}", Method: http.MethodDelete, Handler: HandleDeletePatient},
	{Path: "/patients/{id}", Method: http.MethodPatch, Handler: HandlerUpdatePatient},
	{Path: "/patients/{id}/healthinsurence", Method: http.MethodPost, Handler: HandlerAddHealthInsurenceInPatient},
	{Path: "/doctors", Method: http.MethodPost, Handler: HandleCreateDoctor},
	{Path: "/doctors", Method: http.MethodGet, Handler: HandleGetAllDoctors},
	{Path: "/doctors/{id}", Method: http.MethodDelete, Handler: HandleDeleteDoctor},
	{Path: "/doctors/{id}", Method: http.MethodPatch, Handler: HandlerUpdateDoctor},
	{Path: "/doctors/{id}/healthinsurence", Method: http.MethodPost, Handler: HandlerAddHealthInsurenceInDoctor},
}
