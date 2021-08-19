package service

import (
	"context"
	"github.com/BaytoorJr/kolesa-access/src/transport"
)

type CarDataService interface {
	GetCarsByMark(ctx context.Context, req *transport.GetCarsByMarkRequest) (*transport.GetCarsByMarkResponse, error)
	GetCarsByMarkAndModel(ctx context.Context, req *transport.GetCarsByMarkAndModelRequest) (*transport.GetCarsByMarkAndModelResponse, error)
	GetCarsByYear(ctx context.Context, req *transport.GetCarsByYearRequest) (*transport.GetCarsByYearResponse, error)
	GetCarsNum(ctx context.Context, _ *transport.GetCarsNumRequest) (*transport.GetCarsNumResponse, error)
	GetCarsByAvgPrice(ctx context.Context, req *transport.GetCarsByAvgPriceRequest) (*transport.GetCarsByAvgPriceResponse, error)
	GetCarsByPrice(ctx context.Context, req *transport.GetCarsByPriceRequest) (*transport.GetCarsByPriceResponse, error)
}
