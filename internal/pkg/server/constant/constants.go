package constant

import (
	"github.com/tiagompalte/golang-clean-arch-template/pkg/context"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

var Blank = usecase.Blank{}

const Authorization = "Authorization"

const ContextUser context.ContextKey = "user"
