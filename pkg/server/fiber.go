package server

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/tiagompalte/golang-clean-arch-template/configs"
)

type fiberServer struct {
	app    *fiber.App
	config configs.Config
}

func NewFiberServer(config configs.Config) Server {
	app := fiber.New(fiber.Config{
		AppName:       config.AppName,
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(helmet.New())
	app.Use(cors.New(cors.ConfigDefault))
	app.Get("/swagger/*", swagger.HandlerDefault)

	return fiberServer{
		app:    app,
		config: config,
	}
}

func (s fiberServer) adapterMiddleware(middleware Middleware) fiber.Handler {
	return adaptor.HTTPMiddleware(middleware)
}

func (s fiberServer) appendMiddlewares(middlewares []Middleware) []fiber.Handler {
	ret := make([]fiber.Handler, 0, len(middlewares))
	for _, mw := range middlewares {
		ret = append(ret, s.adapterMiddleware(mw))
	}
	return ret
}

func (s fiberServer) adapterHandler(handler Handler) fiber.Handler {
	return adaptor.HTTPHandler(HandleError(handler))
}

func (s fiberServer) appendHandlers(middlewares []Middleware, handler Handler) []fiber.Handler {
	ret := make([]fiber.Handler, 0, len(middlewares)+1)
	if len(middlewares) > 0 {
		ret = append(ret, s.appendMiddlewares(middlewares)...)
	}
	ret = append(ret, s.adapterHandler(handler))
	return ret
}

func (s fiberServer) appendGroupRoute(groupRoute GroupRoute) func(fiber.Router) {
	return func(router fiber.Router) {
		for _, route := range groupRoute.Routes {
			router.Add(
				route.Method,
				route.Path,
				s.appendHandlers(route.Middlewares, route.Handler)...,
			)
		}

		for _, grpRoute := range groupRoute.GroupRoutes {
			grp := router.Group(grpRoute.Path, s.appendMiddlewares(grpRoute.Middlewares)...)
			grp.Route("", s.appendGroupRoute(grpRoute))
		}
	}
}

func (s fiberServer) RegisterGroupRoutes(groupRoutes []GroupRoute) {
	for _, groupRoute := range groupRoutes {
		grp := s.app.Group(groupRoute.Path, s.appendMiddlewares(groupRoute.Middlewares)...)
		grp.Route("", s.appendGroupRoute(groupRoute))
	}
}

func (s fiberServer) Start() error {
	return s.app.Listen(s.config.WebPort)
}
