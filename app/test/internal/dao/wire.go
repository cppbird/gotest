// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package dao

import (
	"github.com/google/wire"
)

func newTestDao() (*dao, func(), error) {
	panic(wire.Build(newDao, NewDB, NewRedis, NewMC))
}
