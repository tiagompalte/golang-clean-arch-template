package main

import (
	"log"
	"os"

	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server"
)

func main() {
	app, err := application.Build()
	if err != nil {
		log.Fatalf("failed to build the application (error: %v)", err)
	}

	err = server.NewServer(app)
	if err != nil {
		log.Fatalf("failed to start the http server (error: %v)", err)
	}

	os.Exit(0)
}
