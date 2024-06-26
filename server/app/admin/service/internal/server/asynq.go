package server

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/asynq"

	conf "github.com/bugsmo/cy/contrib/kratos/bootstrap/gen/api/go/conf/v1"

	"github.com/bugsmo/chunyu/app/admin/service/internal/service"
)

func NewAsynqServer(cfg *conf.Bootstrap, _ log.Logger, svc *service.TaskService) *asynq.Server {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	srv := asynq.NewServer(
		asynq.WithAddress(cfg.Server.Asynq.GetEndpoint()),
		asynq.WithRedisPassword(cfg.Server.Asynq.GetPassword()),
		asynq.WithRedisDatabase(int(cfg.Server.Asynq.GetDb())),
		asynq.WithLocation(loc.String()),
	)

	svc.Server = srv

	svc.StartAllPeriodicTask()
	svc.StartAllDelayTask()

	return srv
}
