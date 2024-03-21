package api

import (
	"medical-clinic/middleware"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func AdminMiddleware(handler http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.IsAdminMiddleware(db, http.HandlerFunc(handler)).ServeHTTP(w, r)
	}
}

var routes = []Route{
	{Path: "/admin", Method: http.MethodPost, Handler: HandleCreateAdmin},
	{Path: "/admin", Method: http.MethodGet, Handler: HandleGetAllAdmins},
	{Path: "/admins/{id}", Method: http.MethodDelete, Handler: HandleDeleteAdmin},
	{Path: "/admins/{id}", Method: http.MethodPatch, Handler: HandleUpdateAdmin},
	{Path: "/patients", Method: http.MethodPost, Handler: HandleCreatePatient},
	{Path: "/patients", Method: http.MethodGet, Handler: HandleGetAllPatient},
	{Path: "/patients/{id}", Method: http.MethodDelete, Handler: AdminMiddleware(HandleDeletePatient)},
	{Path: "/patients/{id}", Method: http.MethodPatch, Handler: AdminMiddleware(HandlerUpdatePatient)},
	{Path: "/patients/{id}/healthinsurence", Method: http.MethodPost, Handler: HandlerAddHealthInsurenceInPatient},
	{Path: "/doctors", Method: http.MethodPost, Handler: HandleCreateDoctor},
	{Path: "/doctors", Method: http.MethodGet, Handler: HandleGetAllDoctors},
	{Path: "/doctors/{id}", Method: http.MethodDelete, Handler: AdminMiddleware(HandleDeleteDoctor)},
	{Path: "/doctors/{id}", Method: http.MethodPatch, Handler: AdminMiddleware(HandlerUpdateDoctor)},
	{Path: "/doctors/{id}/healthinsurence", Method: http.MethodPost, Handler: HandlerAddHealthInsurenceInDoctor},
}
