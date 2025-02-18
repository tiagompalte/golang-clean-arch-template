package server

type RouteMethod string

const (
	RouteMethodGet     RouteMethod = "GET"
	RouteMethodPost    RouteMethod = "POST"
	RouteMethodPut     RouteMethod = "PUT"
	RouteMethodPatch   RouteMethod = "PATCH"
	RouteMethodDelete  RouteMethod = "DELETE"
	RouteMethodHead    RouteMethod = "HEAD"
	RouteMethodConnect RouteMethod = "CONNECT"
	RouteMethodOptions RouteMethod = "OPTIONS"
	RouteMethodTrace   RouteMethod = "TRACE"
)
