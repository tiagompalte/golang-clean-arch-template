//go:build integration
// +build integration

package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/test/testconfig"
)

func TestMain(t *testing.M) {
	config := testconfig.Instance()

	defer config.Close()

	code := t.Run()

	os.Exit(code)
}

func GenerateUserAndToken() (usecase.CreateUserOutput, string) {
	ctx := context.Background()

	createUserInput := usecase.CreateUserInput{
		Name:     RandomName(),
		Email:    Email(),
		Password: "Pass!1234",
	}

	app := testconfig.Instance().App()

	userLogged, err := app.UseCase().CreateUserUseCase.Execute(ctx, createUserInput)
	if err != nil {
		log.Fatalf("failed to create user logged: %v", err)
	}

	var generateUserTokenInput usecase.GenerateUserTokenInput
	generateUserTokenInput.UUID = userLogged.UUID
	generateUserTokenInput.Name = userLogged.Name
	generateUserTokenInput.Email = userLogged.Email

	token, err := app.UseCase().GenerateUserTokenUseCase.Execute(ctx, generateUserTokenInput)
	if err != nil {
		log.Fatalf("failed to generate token: %v", err)
	}

	return userLogged, token.AccessToken
}
