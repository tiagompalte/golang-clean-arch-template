package crypto

import "context"

type Crypto interface {
	GenerateHash(ctx context.Context, plainText string) (string, error)
	VerifyHash(ctx context.Context, plainText string, hash string) (bool, error)
}
