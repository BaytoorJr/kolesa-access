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
}

// Endpoints middleware constructor
func MakeEndpoints(s service.CarDataService) *Endpoints {
	return &Endpoints{
		GetCarsByMark:         makeGetCarsByMark(s),
		GetCarsByMarkAndModel: makeGetCarsByMarkAndModel(s),
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
