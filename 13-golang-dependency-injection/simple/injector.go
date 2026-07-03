//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService, // memberi tahu ke dalam google wire, function mana saja yang digunakan sebagai provider
	)
	return nil, nil
}
