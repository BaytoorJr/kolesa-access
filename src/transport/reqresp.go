package transport

import "github.com/BaytoorJr/kolesa-access/src/domain"

type (
	// GetCarsByMarkRequest Get car ads by mark req & resp
	GetCarsByMarkRequest struct {
		Mark string `json:"mark"`
	}
	GetCarsByMarkResponse struct {
		Car           []domain.Car `json:"car"`
		RequestsCount int          `json:"requests_count"`
	}

	// GetCarsByMarkAndModelRequest Get car ads by mark and model req & resp
	GetCarsByMarkAndModelRequest struct {
		Mark  string `json:"mark"`
		Model string `json:"model"`
	}
	GetCarsByMarkAndModelResponse struct {
		Car           []domain.Car `json:"car"`
		RequestsCount int          `json:"requests_count"`
	}
)
