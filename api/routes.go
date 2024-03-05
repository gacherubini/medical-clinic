package api

import (
	"net/http"
)

type Route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

var routes = []Route{
	{Path: "/createUser", Handler: HandleCreateUser},
	{Path: "/getUser", Handler: HandleGetAllUser},
	{Path: "/deleteUser/{id}", Handler: HandleDeleteUser},
	{Path: "/updateUser/{id}", Handler: HandleUpdateUser},
}
