package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tiagompalte/golang-clean-arch-template/configs"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type JwtAuth struct {
	algorithm  string
	keySecret  string
	expiration int
}

func NewJwtAuth(config configs.Config) Auth {
	return JwtAuth{
		algorithm:  config.Jwt.Algorithm,
		keySecret:  config.Jwt.KeySecret,
		expiration: config.Jwt.Duration,
	}
}

func (a JwtAuth) GenerateToken(ctx context.Context, mapClaim map[string]any) (string, error) {
	token := jwt.New(jwt.GetSigningMethod(a.algorithm))

	claims := token.Claims.(jwt.MapClaims)
	for key, value := range mapClaim {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(time.Duration(a.expiration * int(time.Second))).Unix()

	jwt, err := token.SignedString([]byte(a.keySecret))
	if err != nil {
		return "", errors.Wrap(err)
	}

	return jwt, nil
}

func (a JwtAuth) ValidateToken(ctx context.Context, token string) (bool, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.keySecret), nil
	})
	if err != nil {
		return false, errors.Wrap(err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return false, nil
	}

	isValid := claims.VerifyExpiresAt(time.Now().Local().Unix(), true)

	return isValid, nil
}

func (a JwtAuth) ValidateExtractToken(ctx context.Context, token string) (map[string]any, bool, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.keySecret), nil
	})
	if err != nil {
		return map[string]any{}, false, errors.Wrap(err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return map[string]any{}, false, nil
	}

	isValid := claims.VerifyExpiresAt(time.Now().Local().Unix(), true)

	ret := make(map[string]any, 0)
	for key, value := range claims {
		ret[key] = value
	}

	return ret, isValid, nil
}
