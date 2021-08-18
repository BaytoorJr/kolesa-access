package postgresql

import (
	"context"
	"git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/errors"
	"github.com/BaytoorJr/kolesa-access/src/config"
	"github.com/BaytoorJr/kolesa-access/src/repository"
	"github.com/go-kit/kit/log"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
)

type Store struct {
	db     *pgxpool.Pool
	logger log.Logger

	CarRepository repository.CarRepository
}

// Store constructor
func NewStore(db *pgxpool.Pool, logger log.Logger) (*Store, error) {
	repo := &Store{
		db:     db,
		logger: log.With(logger, "rep", "postgresql"),
	}

	err := repo.migrate()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// Migration method
func (s *Store) migrate() error {
	for i := 0; i < len(migrations); i++ {
		sql := strings.Replace(migrations[i], "$1", config.MainConfig.PostgreSQLConfig.PostgresSchema, 1)
		_, err := s.db.Exec(context.Background(), sql)
		if err != nil {
			return errors.DBWriteError.SetDevMessage(err.Error())
		}
	}

	return nil
}

func (s *Store) Car() repository.CarRepository {
	if s.CarRepository != nil {
		return s.CarRepository
	}

	s.CarRepository = &CarRepository{
		store: s,
	}

	return s.CarRepository
}
