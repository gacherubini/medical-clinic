package api

import (
	"medical-clinic/admin"
	"medical-clinic/auth"
	"medical-clinic/doctor"
	"medical-clinic/patient"

	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func getDoctorRoutes() []Route {
	DoctorHandlerContext := doctor.DoctorHandlerContext{
		Db: Db,
	}

	var routes = []Route{
		{Path: "/doctors", Method: http.MethodPost, Handler: DoctorHandlerContext.HandleCreateDoctor},
		{Path: "/doctors", Method: http.MethodGet, Handler: DoctorHandlerContext.HandleGetAllDoctors},
		{Path: "/doctors/{id}", Method: http.MethodDelete, Handler: DoctorHandlerContext.HandleDeleteDoctor},
		{Path: "/doctors/{id}", Method: http.MethodPatch, Handler: DoctorHandlerContext.HandlerUpdateDoctor},
		{Path: "/doctors/{id}/healthinsurence", Method: http.MethodPost, Handler: DoctorHandlerContext.HandlerAddHealthInsurenceInDoctor},
	}

	return routes
}

func getAdminRoutes() []Route {
	AdminHandlerContext := admin.AdminHandlerContext{
		Db: Db,
	}
	var routes = []Route{
		{Path: "/admin", Method: http.MethodPost, Handler: AdminHandlerContext.HandleCreateAdmin},
		{Path: "/admin", Method: http.MethodGet, Handler: AdminHandlerContext.HandleGetAllAdmins},
		{Path: "/admins/{id}", Method: http.MethodDelete, Handler: AdminHandlerContext.HandleDeleteAdmin},
		{Path: "/admins/{id}", Method: http.MethodPatch, Handler: AdminHandlerContext.HandleUpdateAdmin},
		{Path: "/admins/healthinsurence", Method: http.MethodPost, Handler: AdminHandlerContext.HandleAdminCreateHealthInsurence},
		{Path: "/admins/healthinsurence", Method: http.MethodGet, Handler: AdminHandlerContext.HandleAdminGetAllHealthInsurence},
	}

	return routes
}

func getPatientRoutes() []Route {
	PatientHandlerContext := patient.PatientHandlerContext{
		Db: Db,
	}
	var routes = []Route{
		{Path: "/patients", Method: http.MethodPost, Handler: PatientHandlerContext.HandleCreatePatient},
		{Path: "/patients", Method: http.MethodGet, Handler: PatientHandlerContext.HandleGetAllPatient},
		{Path: "/patients/{id}", Method: http.MethodDelete, Handler: PatientHandlerContext.HandleDeletePatient},
		{Path: "/patients/{id}", Method: http.MethodPatch, Handler: PatientHandlerContext.HandlerUpdatePatient},
		{Path: "/patients/{id}/healthinsurence", Method: http.MethodPost, Handler: PatientHandlerContext.HandlerAddHealthInsurenceInPatient},
	}

	return routes
}

func getAuthRoutes() []Route {
	AuthHandlerContext := auth.LoginHandlerContext{
		Db: Db,
	}

	var routes = []Route{
		{Path: "/login", Method: http.MethodPost, Handler: AuthHandlerContext.LoginHandler},
	}

	return routes
}
