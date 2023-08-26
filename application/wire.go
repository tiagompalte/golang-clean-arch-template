//go:build wireinject
// +build wireinject

package application

import "github.com/google/wire"

func Build() (App, error) {
	panic(wire.Build(ProviderSet))
}
