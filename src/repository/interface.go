package repository

type MainStore interface {
	Car() CarRepository
}
