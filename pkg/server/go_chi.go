package server

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/tiagompalte/golang-clean-arch-template/configs"
)

type goChiServer struct {
	mux    *chi.Mux
	config configs.Config
}

func NewGoChiServer(config configs.Config) Server {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Mount("/swagger", httpSwagger.WrapHandler)

	return goChiServer{
		mux:    mux,
		config: config,
	}
}

func (s goChiServer) appendMiddlewares(route *chi.Mux, middlewares []Middleware) {
	for _, mw := range middlewares {
		route.Use(mw)
	}
}

func (s goChiServer) adapterHandler(handler Handler) http.HandlerFunc {
	return HandleError(handler)
}

func (s goChiServer) appendRoutes(routeMain *chi.Mux, routes []Route) {
	for _, r := range routes {
		routeMain.Group(func(route chi.Router) {
			for _, mw := range r.Middlewares {
				route.Use(mw)
			}

			switch r.Method {
			case RouteMethodGet:
				route.Get(r.Path, s.adapterHandler(r.Handler))
			case RouteMethodPost:
				route.Post(r.Path, s.adapterHandler(r.Handler))
			case RouteMethodPut:
				route.Put(r.Path, s.adapterHandler(r.Handler))
			case RouteMethodPatch:
				route.Patch(r.Path, s.adapterHandler(r.Handler))
			case RouteMethodDelete:
				route.Delete(r.Path, s.adapterHandler(r.Handler))
			case RouteMethodHead:
				route.Head(r.Path, s.adapterHandler(r.Handler))
			case RouteMethodConnect:
				route.Connect(r.Path, s.adapterHandler(r.Handler))
			case RouteMethodOptions:
				route.Options(r.Path, s.adapterHandler(r.Handler))
			case RouteMethodTrace:
				route.Trace(r.Path, s.adapterHandler(r.Handler))
			}
		})
	}
}

func (s goChiServer) appendGroupRoutes(routeMain *chi.Mux, groupRoutes []GroupRoute) {
	for _, groupRoute := range groupRoutes {
		route := chi.NewRouter()
		s.appendMiddlewares(route, groupRoute.Middlewares)
		s.appendRoutes(route, groupRoute.Routes)
		s.appendGroupRoutes(route, groupRoute.GroupRoutes)
		routeMain.Mount(groupRoute.Path, route)
	}
}

func (s goChiServer) NewServer(groupRoutes []GroupRoute) *http.Server {
	s.appendGroupRoutes(s.mux, groupRoutes)
	return &http.Server{Addr: s.config.WebPort, Handler: s.mux}
}

func (s goChiServer) Start(server *http.Server) error {
	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

	return nil
}

func (s goChiServer) StartTest(httpServer *http.Server) *httptest.Server {
	return httptest.NewServer(httpServer.Handler)
}
