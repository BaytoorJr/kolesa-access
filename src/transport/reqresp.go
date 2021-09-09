package transport

import "github.com/BaytoorJr/kolesa-access/src/domain"

type (
	// GetCarsByMarkRequest Get car ads by mark req & resp
	GetCarsByMarkRequest struct {
		Mark string `json:"mark"`
	}
	GetCarsByMarkResponse struct {
		TotalHits int          `json:"total_hits"`
		Car       []domain.Car `json:"car"`
	}

	// GetCarsByMarkAndModelRequest Get car ads by mark and model req & resp
	GetCarsByMarkAndModelRequest struct {
		Mark  string `json:"mark"`
		Model string `json:"model"`
	}
	GetCarsByMarkAndModelResponse struct {
		TotalHits int          `json:"total_hits"`
		Car       []domain.Car `json:"car"`
	}

	// GetCarsByYearRequest Get car adds by year req & resp
	GetCarsByYearRequest struct {
		Year int `json:"year"`
	}
	GetCarsByYearResponse struct {
		TotalHits int          `json:"total_hits"`
		Car       []domain.Car `json:"car"`
	}

	// GetCarsNumRequest Get number of stored car
	GetCarsNumRequest struct {
	}
	GetCarsNumResponse struct {
		Count int `json:"count"`
	}

	// GetCarsByAvgPriceRequest Get car adds by  average price in min/max range req & resp
	GetCarsByAvgPriceRequest struct {
		Min int `json:"min"`
		Max int `json:"max"`
	}
	GetCarsByAvgPriceResponse struct {
		TotalHits int          `json:"total_hits"`
		Car       []domain.Car `json:"car"`
	}

	// GetCarsByPriceRequest Get car adds by price  in min/max range req & resp
	GetCarsByPriceRequest struct {
		Min int `json:"min"`
		Max int `json:"max"`
	}
	GetCarsByPriceResponse struct {
		TotalHits int          `json:"total_hits"`
		Car       []domain.Car `json:"car"`
	}

	// GetAvgPriceRequest Get average price by mark, model, year
	GetAvgPriceRequest struct {
		Mark  string `json:"mark"`
		Model string `json:"model"`
		Year  int    `json:"year"`
	}
	GetAvgPriceResponse struct {
		AveragePrice int `json:"average_price"`
	}
)
