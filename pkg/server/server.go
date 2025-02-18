package server

import (
	"net/http"
	"net/http/httptest"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

type Middleware func(next http.Handler) http.Handler

type GroupRoute struct {
	Path        string
	GroupRoutes []GroupRoute
	Middlewares []Middleware
	Routes      []Route
}

type Route struct {
	Method      RouteMethod
	Path        string
	Middlewares []Middleware
	Handler     Handler
}

type Server interface {
	NewServer(groupRoutes []GroupRoute) *http.Server
	Start(httpServer *http.Server) error
	StartTest(httpServer *http.Server) *httptest.Server
}
