package testconfig

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server"
)

var config *testConfig

type testConfig struct {
	app        application.App
	httpServer *http.Server
	httpTest   *httptest.Server
}

func newTestConfig() testConfig {
	app, err := application.Build()
	if err != nil {
		log.Fatalf("failed to build the application: %v", err)
	}

	httpServer := server.NewServer(app)
	httpTest := app.Server().StartTest(httpServer)

	return testConfig{
		app,
		httpServer,
		httpTest,
	}
}

func (t *testConfig) App() application.App {
	return t.app
}

func (t *testConfig) HttpUrl() string {
	return t.httpTest.URL
}

func (t *testConfig) Close() {
	t.httpServer.Close()
	t.httpTest.Close()
}

func Instance() *testConfig {
	if config == nil {
		c := newTestConfig()
		config = &c
	}
	return config
}
