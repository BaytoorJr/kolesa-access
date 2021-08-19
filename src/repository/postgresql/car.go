package postgresql

import (
	"context"
	"git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/errors"
	"github.com/BaytoorJr/kolesa-access/src/domain"
)

type CarRepository struct {
	store *Store
}

func (c *CarRepository) GetCarsByMark(ctx context.Context, mark string) (*[]domain.Car, error) {
	var cars []domain.Car

	rows, err := c.store.db.Query(ctx, "SELECT * FROM cars_an2 "+
		"WHERE car_mark = $1", mark)
	if err != nil {
		return nil, errors.DBReadError.SetDevMessage(err.Error())
	}

	for rows.Next() {
		var car domain.Car

		err = rows.Scan(
			&car.ID,
			&car.Link,
			&car.Title,
			&car.ShortDescription,
			&car.Price,
			&car.PostDate,
			&car.ViewsCount,
			&car.AvgPrice,
			&car.City,
			&car.CarMark,
			&car.CarModel,
			&car.CarYear,
			&car.CarType,
			&car.EngineVolume,
			&car.Transmission,
			&car.WheelLocation,
			&car.Color,
			&car.WheelDrive,
			&car.IsKz,
			&car.FullDescription,
			&car.Promotions,
			&car.PromotionsPrice,
			&car.CreatedAt,
			&car.UpdatedAt)

		if err != nil {
			return nil, errors.DBReadError.SetDevMessage(err.Error())
		}

		cars = append(cars, car)

	}

	return &cars, nil
}

func (c *CarRepository) GetCarsByMarkAndModel(ctx context.Context, mark, model string) (*[]domain.Car, error) {
	var cars []domain.Car

	rows, err := c.store.db.Query(ctx, "SELECT * FROM cars_an2 "+
		"WHERE car_mark = $1 and car_model = $2", mark, model)
	if err != nil {
		return nil, errors.DBReadError.SetDevMessage(err.Error())
	}

	for rows.Next() {
		var car domain.Car

		err = rows.Scan(
			&car.ID,
			&car.Link,
			&car.Title,
			&car.ShortDescription,
			&car.Price,
			&car.PostDate,
			&car.ViewsCount,
			&car.AvgPrice,
			&car.City,
			&car.CarMark,
			&car.CarModel,
			&car.CarYear,
			&car.CarType,
			&car.EngineVolume,
			&car.Transmission,
			&car.WheelLocation,
			&car.Color,
			&car.WheelDrive,
			&car.IsKz,
			&car.FullDescription,
			&car.Promotions,
			&car.PromotionsPrice,
			&car.CreatedAt,
			&car.UpdatedAt)

		if err != nil {
			return nil, errors.DBReadError.SetDevMessage(err.Error())
		}

		cars = append(cars, car)

	}

	return &cars, nil
}

func (c *CarRepository) GetCarsByYear(ctx context.Context, year int) (*[]domain.Car, error) {
	var cars []domain.Car

	rows, err := c.store.db.Query(ctx, "SELECT * FROM cars_an2 "+
		"WHERE car_year = $1", year)
	if err != nil {
		return nil, errors.DBReadError.SetDevMessage(err.Error())
	}

	for rows.Next() {
		var car domain.Car

		err = rows.Scan(
			&car.ID,
			&car.Link,
			&car.Title,
			&car.ShortDescription,
			&car.Price,
			&car.PostDate,
			&car.ViewsCount,
			&car.AvgPrice,
			&car.City,
			&car.CarMark,
			&car.CarModel,
			&car.CarYear,
			&car.CarType,
			&car.EngineVolume,
			&car.Transmission,
			&car.WheelLocation,
			&car.Color,
			&car.WheelDrive,
			&car.IsKz,
			&car.FullDescription,
			&car.Promotions,
			&car.PromotionsPrice,
			&car.CreatedAt,
			&car.UpdatedAt)

		if err != nil {
			return nil, errors.DBReadError.SetDevMessage(err.Error())
		}

		cars = append(cars, car)
	}

	return &cars, nil
}

func (c *CarRepository) GetRowsNum(ctx context.Context) (int, error) {
	var count int

	err := c.store.db.QueryRow(ctx, "SELECT COUNT(*) FROM cars_an2").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c *CarRepository) GetCarsByAvgPrice(ctx context.Context, min, max int) (*[]domain.Car, error) {
	var cars []domain.Car

	rows, err := c.store.db.Query(ctx, "SELECT * FROM cars_an2 "+
		"WHERE average_price >= $1 and average_price <= $2", min, max)
	if err != nil {
		return nil, errors.DBReadError.SetDevMessage(err.Error())
	}

	for rows.Next() {
		var car domain.Car

		err = rows.Scan(
			&car.ID,
			&car.Link,
			&car.Title,
			&car.ShortDescription,
			&car.Price,
			&car.PostDate,
			&car.ViewsCount,
			&car.AvgPrice,
			&car.City,
			&car.CarMark,
			&car.CarModel,
			&car.CarYear,
			&car.CarType,
			&car.EngineVolume,
			&car.Transmission,
			&car.WheelLocation,
			&car.Color,
			&car.WheelDrive,
			&car.IsKz,
			&car.FullDescription,
			&car.Promotions,
			&car.PromotionsPrice,
			&car.CreatedAt,
			&car.UpdatedAt)

		if err != nil {
			return nil, errors.DBReadError.SetDevMessage(err.Error())
		}

		cars = append(cars, car)
	}

	return &cars, nil

}

func (c *CarRepository) GetCarsByPrice(ctx context.Context, min, max int) (*[]domain.Car, error) {

	var cars []domain.Car

	rows, err := c.store.db.Query(ctx, "SELECT * FROM cars_an2 "+
		"WHERE price >= $1 and price <= $2", min, max)
	if err != nil {
		return nil, errors.DBReadError.SetDevMessage(err.Error())
	}

	for rows.Next() {
		var car domain.Car

		err = rows.Scan(
			&car.ID,
			&car.Link,
			&car.Title,
			&car.ShortDescription,
			&car.Price,
			&car.PostDate,
			&car.ViewsCount,
			&car.AvgPrice,
			&car.City,
			&car.CarMark,
			&car.CarModel,
			&car.CarYear,
			&car.CarType,
			&car.EngineVolume,
			&car.Transmission,
			&car.WheelLocation,
			&car.Color,
			&car.WheelDrive,
			&car.IsKz,
			&car.FullDescription,
			&car.Promotions,
			&car.PromotionsPrice,
			&car.CreatedAt,
			&car.UpdatedAt)

		if err != nil {
			return nil, errors.DBReadError.SetDevMessage(err.Error())
		}

		cars = append(cars, car)
	}

	return &cars, nil

}
