package middleware

import (
	"context"
	"github.com/BaytoorJr/kolesa-access/src/service"
	"github.com/BaytoorJr/kolesa-access/src/transport"
	"github.com/go-kit/kit/metrics"
	"time"
)

// Instrumenting middleware entity
type instrumentingMiddleware struct {
	next           service.CarDataService
	requestCount   metrics.Counter
	requestError   metrics.Counter
	requestLatency metrics.Histogram
}

// Instrumenting middleware private method
func (im *instrumentingMiddleware) instrumenting(begin time.Time, method string, err error) {
	im.requestCount.With("method", method).Add(1)
	if err != nil {
		im.requestError.With("method", method).Add(1)
	}
	im.requestLatency.With("method", method).Observe(time.Since(begin).Seconds())
}

// Instrumenting middleware constructor
func NewInstrumentingMiddleware(counter, counterErr metrics.Counter, latency metrics.Histogram) Middleware {
	return func(next service.CarDataService) service.CarDataService {
		return &instrumentingMiddleware{
			next:           next,
			requestCount:   counter,
			requestError:   counterErr,
			requestLatency: latency,
		}
	}
}

func (im *instrumentingMiddleware) GetCarsByMark(ctx context.Context, req *transport.GetCarsByMarkRequest) (_ *transport.GetCarsByMarkResponse, err error) {
	defer im.instrumenting(time.Now(), "GetCarsByMark", err)
	return im.next.GetCarsByMark(ctx, req)
}

func (im *instrumentingMiddleware) GetCarsByMarkAndModel(ctx context.Context, req *transport.GetCarsByMarkAndModelRequest) (_ *transport.GetCarsByMarkAndModelResponse, err error) {
	defer im.instrumenting(time.Now(), "GetCarsByMarkAndModel", err)
	return im.next.GetCarsByMarkAndModel(ctx, req)
}
