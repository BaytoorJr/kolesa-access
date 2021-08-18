package repository

import (
	"context"
	"github.com/BaytoorJr/kolesa-access/src/domain"
)

type CarRepository interface {
	GetCarsByMark(ctx context.Context, mark string) (*[]domain.Car, error)
	GetCarsByMarkAndModel(ctx context.Context, mark, model string) (*[]domain.Car, error)
	GetCarsByYear(ctx context.Context, year int) (*[]domain.Car, error)
}
