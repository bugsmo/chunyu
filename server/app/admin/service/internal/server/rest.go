package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"

	authnEngine "github.com/bugsmo/cy/contrib/kratos/authn/engine"
	authn "github.com/bugsmo/cy/contrib/kratos/authn/middleware"

	authzEngine "github.com/bugsmo/cy/contrib/kratos/authz/engine"
	authz "github.com/bugsmo/cy/contrib/kratos/authz/middleware"

	swaggerUI "github.com/bugsmo/cy/contrib/kratos/swagger-ui"

	bootstrap "github.com/bugsmo/cy/contrib/kratos/bootstrap"
	conf "github.com/bugsmo/cy/contrib/kratos/bootstrap/gen/api/go/conf/v1"

	"github.com/bugsmo/chunyu/app/admin/service/cmd/server/assets"
	"github.com/bugsmo/chunyu/app/admin/service/internal/service"

	adminV1 "github.com/bugsmo/chunyu/gen/api/go/admin/service/v1"
	"github.com/bugsmo/chunyu/pkg/middleware/auth"
)

// NewWhiteListMatcher 创建jwt白名单
func newRestWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList[adminV1.OperationAuthenticationServiceLogin] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddleware 创建中间件
func newRestMiddleware(authenticator authnEngine.Authenticator, authorizer authzEngine.Engine, logger log.Logger) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))
	ms = append(ms, selector.Server(
		authn.Server(authenticator),
		auth.Server(),
		authz.Server(authorizer),
	).Match(newRestWhiteListMatcher()).Build())
	return ms
}

// NewRESTServer new an HTTP server.
func NewRESTServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authnEngine.Authenticator, authorizer authzEngine.Engine,
	authnSvc *service.AuthenticationService,
	userSvc *service.UserService,
	roleSvc *service.RoleService,
	// dictSvc *service.DictService,
	// dictDetailSvc *service.DictDetailService,
	// menuSvc *service.MenuService,
	// routerSvc *service.RouterService,
	// orgSvc *service.OrganizationService,
	// positionSvc *service.PositionService,
) *http.Server {
	srv := bootstrap.CreateRestServer(cfg, newRestMiddleware(authenticator, authorizer, logger)...)

	adminV1.RegisterAuthenticationServiceHTTPServer(srv, authnSvc)
	adminV1.RegisterUserServiceHTTPServer(srv, userSvc)
	// adminV1.RegisterDictServiceHTTPServer(srv, dictSvc)
	// adminV1.RegisterDictDetailServiceHTTPServer(srv, dictDetailSvc)
	// adminV1.RegisterMenuServiceHTTPServer(srv, menuSvc)
	// adminV1.RegisterRouterServiceHTTPServer(srv, routerSvc)
	// adminV1.RegisterOrganizationServiceHTTPServer(srv, orgSvc)
	adminV1.RegisterRoleServiceHTTPServer(srv, roleSvc)
	// adminV1.RegisterPositionServiceHTTPServer(srv, positionSvc)

	if cfg.GetServer().GetRest().GetEnableSwagger() {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithTitle("Kratos巨石应用实践"),
			swaggerUI.WithMemoryData(assets.OpenApiData, "yaml"),
			swaggerUI.WithBasePath("/docs/"),
		)
	}

	return srv
}
