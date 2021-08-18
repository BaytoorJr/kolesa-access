package constructor

import (
	"context"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/BaytoorJr/kolesa-access/src/middleware"
	"github.com/BaytoorJr/kolesa-access/src/repository"
	"github.com/BaytoorJr/kolesa-access/src/service"
)

// Car data service constructor
func NewCarDataService(ctx context.Context, mainRepo repository.MainStore, logger log.Logger) service.CarDataService {
	svc := service.NewService(ctx, mainRepo, logger)
	svc = middleware.NewLoggingMiddleware(logger)(svc)
	svc = middleware.NewInstrumentingMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "car_data_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "car_data_service",
			Name:      "error_count",
			Help:      "Number of error requests received.",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "car_data_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}),
	)(svc)

	return svc
}
