package usecase

import (
	"context"
	"testing"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/auth"
)

func TestNewGenerateUserTokenExecute(t *testing.T) {
	t.Parallel()

	t.Run("should be return token", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		auth := auth.NewAuthMock("token", map[string]any{})

		us := GenerateUserTokenUseCaseImpl{
			auth: auth,
		}

		result, err := us.Execute(ctx, GenerateUserTokenInput{})

		if err != nil {
			t.Error(err)
		}

		if result.AccessToken != "token" {
			t.Errorf(`token should be "token" but is %s`, result.AccessToken)
		}
	})
}
