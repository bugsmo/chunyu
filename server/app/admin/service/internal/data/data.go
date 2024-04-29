package data

import (
	"github.com/bugsmo/chunyu/app/admin/service/internal/data/ent"
	authnEngine "github.com/bugsmo/cy/contrib/kratos/authn/engine"
	"github.com/bugsmo/cy/contrib/kratos/authn/engine/jwt"
	authzEngine "github.com/bugsmo/cy/contrib/kratos/authz/engine"
	"github.com/bugsmo/cy/contrib/kratos/authz/engine/noop"
	"github.com/bugsmo/cy/contrib/kratos/bootstrap"
	conf "github.com/bugsmo/cy/contrib/kratos/bootstrap/gen/api/go/conf/v1"
	"github.com/bugsmo/cy/contrib/orm/entgo"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/redis/go-redis/v9"
)

type Data struct {
	log           *log.Helper
	rdb           *redis.Client
	db            *entgo.EntClient[*ent.Client]
	authenticator authnEngine.Authenticator
	authorizer    authzEngine.Engine
}

func NewData(
	entClient *entgo.EntClient[*ent.Client],
	redisClient *redis.Client,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Engine,
	logger log.Logger,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/admin-service"))
	d := &Data{
		db:            entClient,
		rdb:           redisClient,
		authenticator: authenticator,
		authorizer:    authorizer,
		log:           l,
	}
	return d, func() {
		l.Info("message", "closing the data resource")
		d.db.Close()
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

func NewRedisClient(cfg *conf.Bootstrap) *redis.Client {
	return bootstrap.NewRedisClient(cfg.GetData())
}

func NewAuthenticator(cfg *conf.Bootstrap) authnEngine.Authenticator {
	authenticator, _ := jwt.NewAuthenticator(
		jwt.WithKey([]byte(cfg.Server.Rest.Middleware.Auth.Key)),
		jwt.WithSigningMethod(cfg.Server.Rest.Middleware.Auth.Method),
	)
	return authenticator
}

func NewAuthorizer() authzEngine.Engine {
	return noop.State{}
}
