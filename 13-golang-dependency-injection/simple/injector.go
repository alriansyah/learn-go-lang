//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

// Function Injector
func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService, // memberi tahu ke dalam google wire, function mana saja yang digunakan sebagai provider
	)
	return nil, nil
}

// Function Injector
func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

// Function Injector
func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}
