package data

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/bugsmo/chunyu/app/admin/service/internal/data/ent"
	"github.com/bugsmo/chunyu/app/admin/service/internal/data/ent/migrate"
	conf "github.com/bugsmo/cy/contrib/kratos/bootstrap/gen/api/go/conf/v1"
	"github.com/bugsmo/cy/contrib/orm/entgo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewEntClient(cfg *conf.Bootstrap, logger log.Logger) *entgo.EntClient[*ent.Client] {
	l := log.NewHelper(log.With(logger, "module", "ent/data/admin-service"))

	drv, err := entgo.CreateDriver(cfg.Data.Database.Driver,
		cfg.Data.Database.Source,
		int(cfg.Data.Database.MaxIdleConnections),
		int(cfg.Data.Database.MaxOpenConnections),
		cfg.Data.Database.ConnectionMaxLifetime.AsDuration(),
	)
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
		return nil
	}

	client := ent.NewClient(ent.Driver(drv),
		ent.Log(func(a ...any) {
			l.Debug(a...)
		}),
	)

	// 运行数据库迁移工具
	if cfg.Data.Database.Migrate {
		if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(true)); err != nil {
			l.Fatalf("failed creating schema resources: %v", err)
		}
	}

	return entgo.NewEntClient(client, drv)
}
