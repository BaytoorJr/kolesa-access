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
		TotalHits: len(*cars),
		Car:       *cars,
	}, nil
}

func (s *service) GetCarsByMarkAndModel(ctx context.Context, req *transport.GetCarsByMarkAndModelRequest) (*transport.GetCarsByMarkAndModelResponse, error) {

	cars, err := s.mainStore.Car().GetCarsByMarkAndModel(ctx, req.Mark, req.Model)
	if err != nil {
		return nil, err
	}

	return &transport.GetCarsByMarkAndModelResponse{
		TotalHits: len(*cars),
		Car:       *cars,
	}, nil
}

func (s *service) GetCarsByYear(ctx context.Context, req *transport.GetCarsByYearRequest) (*transport.GetCarsByYearResponse, error) {
	cars, err := s.mainStore.Car().GetCarsByYear(ctx, req.Year)
	if err != nil {
		return nil, err
	}

	return &transport.GetCarsByYearResponse{
		TotalHits: len(*cars),
		Car:       *cars,
	}, nil
}

func (s *service) GetCarsNum(ctx context.Context, _ *transport.GetCarsNumRequest) (*transport.GetCarsNumResponse, error) {
	count, err := s.mainStore.Car().GetRowsNum(ctx)
	if err != nil {
		return nil, err
	}

	return &transport.GetCarsNumResponse{
		Count: count,
	}, nil
}

func (s *service) GetCarsByAvgPrice(ctx context.Context, req *transport.GetCarsByAvgPriceRequest) (*transport.GetCarsByAvgPriceResponse, error) {
	cars, err := s.mainStore.Car().GetCarsByAvgPrice(ctx, req.Min, req.Max)
	if err != nil {
		return nil, err
	}

	return &transport.GetCarsByAvgPriceResponse{
		TotalHits: len(*cars),
		Car:       *cars,
	}, nil
}

func (s *service) GetCarsByPrice(ctx context.Context, req *transport.GetCarsByPriceRequest) (*transport.GetCarsByPriceResponse, error) {
	cars, err := s.mainStore.Car().GetCarsByPrice(ctx, req.Min, req.Max)
	if err != nil {
		return nil, err
	}

	return &transport.GetCarsByPriceResponse{
		TotalHits: len(*cars),
		Car:       *cars,
	}, nil
}
