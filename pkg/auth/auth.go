package auth

import "context"

type Auth interface {
	GenerateToken(ctx context.Context, mapClaim map[string]any) (string, error)
	ValidateToken(ctx context.Context, token string) (bool, error)
	ValidateExtractToken(ctx context.Context, token string) (map[string]any, bool, error)
}
