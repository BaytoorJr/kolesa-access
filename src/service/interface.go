package service

import (
	"context"
	"github.com/BaytoorJr/kolesa-access/src/transport"
)

type CarDataService interface {
	GetCarsByMark(ctx context.Context, req *transport.GetCarsByMarkRequest) (*transport.GetCarsByMarkResponse, error)
	GetCarsByMarkAndModel(ctx context.Context, req *transport.GetCarsByMarkAndModelRequest) (*transport.GetCarsByMarkAndModelResponse, error)
}
