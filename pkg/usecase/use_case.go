package usecase

import "context"

type Blank struct{}

type UseCase[T any, U any] interface {
	Execute(context context.Context, input T) (U, error)
}
