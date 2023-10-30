package crypto

import "context"

type CryptoMock struct{}

func NewCryptoMock() Crypto {
	return CryptoMock{}
}

func (c CryptoMock) GenerateHash(ctx context.Context, plainText string) (string, error) {
	return plainText, nil
}

func (c CryptoMock) VerifyHash(ctx context.Context, plainText string, hash string) (bool, error) {
	return plainText == hash, nil
}
