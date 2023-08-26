package server

import "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"

var HttpStatusCode = map[string]int{
	errors.ErrorCodeEmptyParameter: 400,
}
