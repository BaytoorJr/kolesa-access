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

// Logging middleware constructor
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
