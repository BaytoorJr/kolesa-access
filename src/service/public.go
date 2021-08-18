package service

import (
	"context"
	"fmt"
	"github.com/BaytoorJr/kolesa-access/src/transport"
)

func (s *service) GetCarsByMark(ctx context.Context, req *transport.GetCarsByMarkRequest) (*transport.GetCarsByMarkResponse, error) {

	cars, err := s.mainStore.Car().GetCarsByMark(ctx, req.Mark)
	if err != nil {
		return nil, err
	}

	fmt.Println(*req)

	return &transport.GetCarsByMarkResponse{
		Car:           *cars,
		RequestsCount: len(*cars),
	}, nil
}

func (s *service) GetCarsByMarkAndModel(ctx context.Context, req *transport.GetCarsByMarkAndModelRequest) (*transport.GetCarsByMarkAndModelResponse, error) {

	cars, err := s.mainStore.Car().GetCarsByMarkAndModel(ctx, req.Mark, req.Model)
	if err != nil {
		return nil, err
	}

	return &transport.GetCarsByMarkAndModelResponse{
		Car:           *cars,
		RequestsCount: len(*cars),
	}, nil
}
