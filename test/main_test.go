//go:build integration
// +build integration

package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/config"
)

var httpTestUrl = ""
var bearerToken = ""

func TestMain(t *testing.M) {
	ctx := context.Background()

	os.Setenv(config.Env, config.EnvTest)

	app, err := application.Build()
	if err != nil {
		log.Fatalf("failed to build the application: %v", err)
	}

	httpServer := server.NewServer(app)
	httpTest := app.Server().StartTest(httpServer)

	httpTestUrl = httpTest.URL

	defer httpTest.Close()
	defer httpServer.Close()

	createUserInput := usecase.CreateUserInput{
		Name:     "User Logged",
		Email:    "user_logged@email.com",
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
	bearerToken = token.AccessToken

	code := t.Run()

	fmt.Println(code)
}
