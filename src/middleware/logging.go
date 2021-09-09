package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/BaytoorJr/kolesa-access/src/service"
	"github.com/BaytoorJr/kolesa-access/src/transport"
)

// Logging middleware entity
type loggingMiddleware struct {
	next   service.CarDataService
	logger log.Logger
}

// Logging middleware private method
func (lm *loggingMiddleware) logging(begin time.Time, method string, err error) {
	_ = lm.logger.Log("method", method, "took", time.Since(begin), "err", err)
}

// NewLoggingMiddleware constructor
func NewLoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.CarDataService) service.CarDataService {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (lm *loggingMiddleware) GetCarsByMark(ctx context.Context, req *transport.GetCarsByMarkRequest) (_ *transport.GetCarsByMarkResponse, err error) {
	defer lm.logging(time.Now(), "GetCarsByMark", err)
	return lm.next.GetCarsByMark(ctx, req)
}

func (lm *loggingMiddleware) GetCarsByMarkAndModel(ctx context.Context, req *transport.GetCarsByMarkAndModelRequest) (_ *transport.GetCarsByMarkAndModelResponse, err error) {
	defer lm.logging(time.Now(), "GetCarsByMarkAndModel", err)
	return lm.next.GetCarsByMarkAndModel(ctx, req)
}

func (lm *loggingMiddleware) GetCarsByYear(ctx context.Context, req *transport.GetCarsByYearRequest) (_ *transport.GetCarsByYearResponse, err error) {
	defer lm.logging(time.Now(), "GetCarsByYear", err)
	return lm.next.GetCarsByYear(ctx, req)
}

func (lm *loggingMiddleware) GetCarsNum(ctx context.Context, req *transport.GetCarsNumRequest) (_ *transport.GetCarsNumResponse, err error) {
	defer lm.logging(time.Now(), "GetCarsNum", err)
	return lm.next.GetCarsNum(ctx, req)
}

func (lm *loggingMiddleware) GetCarsByAvgPrice(ctx context.Context, req *transport.GetCarsByAvgPriceRequest) (_ *transport.GetCarsByAvgPriceResponse, err error) {
	defer lm.logging(time.Now(), "GetCarsByAvgPrice", err)
	return lm.next.GetCarsByAvgPrice(ctx, req)
}

func (lm *loggingMiddleware) GetCarsByPrice(ctx context.Context, req *transport.GetCarsByPriceRequest) (_ *transport.GetCarsByPriceResponse, err error) {
	defer lm.logging(time.Now(), "GetCarsByPrice", err)
	return lm.next.GetCarsByPrice(ctx, req)
}

func (lm *loggingMiddleware) GetAvgPrice(ctx context.Context, req *transport.GetAvgPriceRequest) (_ *transport.GetAvgPriceResponse, err error) {
	defer lm.logging(time.Now(), "GetAveragePrice", err)
	return lm.next.GetAvgPrice(ctx, req)
}
