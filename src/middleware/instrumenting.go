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

// NewInstrumentingMiddleware constructor
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

func (im *instrumentingMiddleware) GetCarsByYear(ctx context.Context, req *transport.GetCarsByYearRequest) (_ *transport.GetCarsByYearResponse, err error) {
	defer im.instrumenting(time.Now(), "GetCarsByYear", err)
	return im.next.GetCarsByYear(ctx, req)
}

func (im *instrumentingMiddleware) GetCarsNum(ctx context.Context, req *transport.GetCarsNumRequest) (_ *transport.GetCarsNumResponse, err error) {
	defer im.instrumenting(time.Now(), "GetCarsNum", err)
	return im.next.GetCarsNum(ctx, req)
}

func (im *instrumentingMiddleware) GetCarsByAvgPrice(ctx context.Context, req *transport.GetCarsByAvgPriceRequest) (_ *transport.GetCarsByAvgPriceResponse, err error) {
	defer im.instrumenting(time.Now(), "GetCarsByAvgPrice", err)
	return im.next.GetCarsByAvgPrice(ctx, req)
}

func (im *instrumentingMiddleware) GetCarsByPrice(ctx context.Context, req *transport.GetCarsByPriceRequest) (_ *transport.GetCarsByPriceResponse, err error) {
	defer im.instrumenting(time.Now(), "GetCarsByPrice", err)
	return im.next.GetCarsByPrice(ctx, req)
}

func (im *instrumentingMiddleware) GetAvgPrice(ctx context.Context, req *transport.GetAvgPriceRequest) (_ *transport.GetAvgPriceResponse, err error) {
	defer im.instrumenting(time.Now(), "GetAveragePrice", err)
	return im.next.GetAvgPrice(ctx, req)
}
