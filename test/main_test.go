//go:build integration
// +build integration

package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server"
)

var app application.App
var httpTestUrl = ""

func TestMain(t *testing.M) {
	var err error
	app, err = application.Build()
	if err != nil {
		log.Fatalf("failed to build the application: %v", err)
	}

	httpServer := server.NewServer(app)
	httpTest := app.Server().StartTest(httpServer)

	httpTestUrl = httpTest.URL

	defer httpTest.Close()
	defer httpServer.Close()

	code := t.Run()

	os.Exit(code)
}

func GenerateUserAndToken() (entity.User, string) {
	ctx := context.Background()

	createUserInput := usecase.CreateUserInput{
		Name:     RandomName(),
		Email:    Email(),
		Password: "Pass!1234",
	}

	userLogged, err := app.UseCase().CreateUserUseCase.Execute(ctx, createUserInput)
	if err != nil {
		log.Fatalf("failed to create user logged: %v", err)
	}

	token, err := app.UseCase().GenerateUserTokenUseCase.Execute(ctx, userLogged)
	if err != nil {
		log.Fatalf("failed to generate token: %v", err)
	}

	return userLogged, token.AccessToken
}
