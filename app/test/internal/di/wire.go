// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"gotest/app/test/internal/dao"
	"gotest/app/test/internal/service"
	"gotest/app/test/internal/server/grpc"
	"gotest/app/test/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
