package server

import (
	// v1 "github.com/mcwmengxi/go/kratos-learn/urlshorter/api/user/v1"
	"github.com/mcwmengxi/go/kratos-learn/urlshorter/internal/conf"
	"github.com/mcwmengxi/go/kratos-learn/urlshorter/internal/interfaces"
	"github.com/mcwmengxi/go/kratos-learn/urlshorter/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,userRouter *interfaces.UserUseCase, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", interfaces.RegisterHTTPServer(userRouter))
	// v1.RegisterGreeterHTTPServer(srv, greeter)
	// v1.RegisterUserHTTPServer(srv, user)
	return srv
}
