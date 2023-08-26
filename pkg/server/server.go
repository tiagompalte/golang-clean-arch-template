package server

import (
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

type Middleware func(h http.Handler) http.Handler

type GroupRoute struct {
	Path        string
	GroupRoutes []GroupRoute
	Middlewares []Middleware
	Routes      []Route
}

type Route struct {
	Method      string
	Path        string
	Middlewares []Middleware
	Handler     Handler
}

type Server interface {
	RegisterGroupRoutes(groupRoutes []GroupRoute)
	Start() error
}
