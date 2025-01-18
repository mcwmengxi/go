//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/mcwmengxi/go/kratos-learn/urlshorter/internal/biz"
	"github.com/mcwmengxi/go/kratos-learn/urlshorter/internal/conf"
	"github.com/mcwmengxi/go/kratos-learn/urlshorter/internal/data"
	"github.com/mcwmengxi/go/kratos-learn/urlshorter/internal/interfaces"
	"github.com/mcwmengxi/go/kratos-learn/urlshorter/internal/server"
	"github.com/mcwmengxi/go/kratos-learn/urlshorter/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, interfaces.ProviderSet , newApp))
}
