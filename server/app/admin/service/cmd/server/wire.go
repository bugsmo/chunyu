//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	conf "github.com/bugsmo/cy/contrib/kratos/bootstrap/gen/api/go/conf/v1"

	"github.com/bugsmo/chunyu/app/admin/service/internal/data"
	"github.com/bugsmo/chunyu/app/admin/service/internal/server"
	"github.com/bugsmo/chunyu/app/admin/service/internal/service"
)

// initApp init kratos application.
func initApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, data.ProviderSet, newApp))
}
