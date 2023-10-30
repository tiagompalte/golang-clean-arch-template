package auth

import "context"

type AuthMock struct {
	token    string
	mapClaim map[string]any
}

func NewAuthMock(token string, mapClaim map[string]any) Auth {
	return AuthMock{
		token:    token,
		mapClaim: mapClaim,
	}
}

func (a AuthMock) GenerateToken(ctx context.Context, mapClaim map[string]any) (string, error) {
	return a.token, nil
}

func (a AuthMock) ValidateToken(ctx context.Context, token string) (bool, error) {
	return a.token == token, nil
}

func (a AuthMock) ValidateExtractToken(ctx context.Context, token string) (map[string]any, bool, error) {
	return a.mapClaim, a.token == token, nil
}
