package app

import (
	"github.com/adexcell/event-booker/internal/entity/adapter/postgres"
	"github.com/adexcell/event-booker/internal/entity/adapter/rabbit"
	"github.com/adexcell/event-booker/internal/entity/adapter/redis"
	httprouter "github.com/adexcell/event-booker/internal/entity/controller/http_router"
	"github.com/adexcell/event-booker/internal/entity/usecase"
)

func EntityDomain(d Dependencies) {
	entityUseCase := usecase.New(
		postgres.New(d.Postgres),
		redis.New(d.Redis),
		rabbit.New(d.RabbitMQ),
	)

	httprouter.EntityRouter(d.RouterHTTP, entityUseCase, d.Metrics)
}
