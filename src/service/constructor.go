package service

import (
	"context"
	"github.com/BaytoorJr/kolesa-access/src/repository"
	"github.com/go-kit/kit/log"
)

// Service entity
type service struct {
	mainStore repository.MainStore
	logger    log.Logger
	context   context.Context
}

// NewService constructor
func NewService(ctx context.Context, mainStore repository.MainStore, logger log.Logger) CarDataService {
	return &service{
		mainStore: mainStore,
		logger:    logger,
		context:   ctx,
	}
}
