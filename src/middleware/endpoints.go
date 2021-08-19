package middleware

import (
	"context"
	"github.com/BaytoorJr/kolesa-access/src/service"
	"github.com/BaytoorJr/kolesa-access/src/transport"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints middleware entity
type Endpoints struct {
	GetCarsByMark         endpoint.Endpoint
	GetCarsByMarkAndModel endpoint.Endpoint
	GetCarsByYear         endpoint.Endpoint
	GetCarsNum            endpoint.Endpoint
	GetCarsByAvgPrice     endpoint.Endpoint
	GetCarsByPrice        endpoint.Endpoint
}

// Endpoints middleware constructor
func MakeEndpoints(s service.CarDataService) *Endpoints {
	return &Endpoints{
		GetCarsByMark:         makeGetCarsByMark(s),
		GetCarsByMarkAndModel: makeGetCarsByMarkAndModel(s),
		GetCarsByYear:         makeGetCarsByYear(s),
		GetCarsNum:            makeGetCarsNum(s),
		GetCarsByAvgPrice:     makerGetCarsByAvgNum(s),
		GetCarsByPrice:        makerGetCarsByNum(s),
	}
}

// Get car adds by mark endpoint
func makeGetCarsByMark(s service.CarDataService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.GetCarsByMarkRequest)
		return s.GetCarsByMark(ctx, &req)
	}
}

// Get car adds by mark & model endpoint
func makeGetCarsByMarkAndModel(s service.CarDataService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.GetCarsByMarkAndModelRequest)
		return s.GetCarsByMarkAndModel(ctx, &req)
	}
}

// Get car adds by year endpoint
func makeGetCarsByYear(s service.CarDataService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.GetCarsByYearRequest)
		return s.GetCarsByYear(ctx, &req)
	}
}

// Count car adds
func makeGetCarsNum(s service.CarDataService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.GetCarsNumRequest)
		return s.GetCarsNum(ctx, &req)
	}
}

// Get car adds by average price in range min and max
func makerGetCarsByAvgNum(s service.CarDataService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.GetCarsByAvgPriceRequest)
		return s.GetCarsByAvgPrice(ctx, &req)
	}
}

// Get car adds by price in range min and max
func makerGetCarsByNum(s service.CarDataService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.GetCarsByPriceRequest)
		return s.GetCarsByPrice(ctx, &req)
	}
}
