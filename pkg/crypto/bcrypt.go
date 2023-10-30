package crypto

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/configs"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	round int
}

func NewBcrypt(config configs.ConfigBcrypt) Crypto {
	return Bcrypt{
		round: config.Round,
	}
}

func (c Bcrypt) GenerateHash(ctx context.Context, plainText string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), c.round)
	if err != nil {
		return "", errors.Wrap(err)
	}
	return string(bytes), err
}

func (c Bcrypt) VerifyHash(ctx context.Context, plainText string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText))
	return err == nil, nil
}
