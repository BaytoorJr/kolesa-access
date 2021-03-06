package http

import (
	encoders "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/transport/http"
	"github.com/BaytoorJr/kolesa-access/src/middleware"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func initializeRoutes(endpoints *middleware.Endpoints, options []kithttp.ServerOption) *mux.Router {
	getCarsByMark := kithttp.NewServer(
		endpoints.GetCarsByMark,
		getCarsByMarkDecoders,
		encoders.EncodeResponse,
		options...,
	)

	getCarsByMarkAndModel := kithttp.NewServer(
		endpoints.GetCarsByMarkAndModel,
		getCarsByMarkAndModelDecoders,
		encoders.EncodeResponse,
		options...,
	)

	getCarsByYear := kithttp.NewServer(
		endpoints.GetCarsByYear,
		getCarsByYearDecoders,
		encoders.EncodeResponse,
		options...,
	)

	getCarsNum := kithttp.NewServer(
		endpoints.GetCarsNum,
		getCarsNumDecoders,
		encoders.EncodeResponse,
		options...,
	)

	getCarsByAvgPrice := kithttp.NewServer(
		endpoints.GetCarsByAvgPrice,
		getCarsByAvgPriceDecoders,
		encoders.EncodeResponse,
		options...,
	)

	getCarsByPrice := kithttp.NewServer(
		endpoints.GetCarsByPrice,
		getCarsByPriceDecoders,
		encoders.EncodeResponse,
		options...,
	)

	getAvgPrice := kithttp.NewServer(
		endpoints.GetAvgPrice,
		getAvgPriceDecoders,
		encoders.EncodeResponse,
		options...,
	)

	router := mux.NewRouter()

	// swagger:route GET /parser-api/kolesa/mark Cars GetCarsByMarkRequest
	// Get car adds by mark
	// responses:
	//   200: GetCarsByMarkResponse
	router.Path("/parser-api/kolesa/car/mark").
		Methods("GET").
		Handler(getCarsByMark)

	// swagger:route GET /parser-api/kolesa/mark_model Cars GetCarsByMarkAndModelRequest
	// Get car adds by mark and model
	// responses:
	//   200: GetCarsByMarkAndModelResponse
	router.Path("/parser-api/kolesa/car/mark&model").
		Methods("GET").
		Handler(getCarsByMarkAndModel)

	// swagger:route GET /parser-api/kolesa/year Cars GetCarsByYearRequest
	// Get car adds by year
	// responses:
	//   200: GetCarsByYearResponse
	router.Path("/parser-api/kolesa/car/year").
		Methods("GET").
		Handler(getCarsByYear)

	// swagger:route GET /parser-api/kolesa/count Cars GetCarsNumRequest
	// Get car adds count
	// responses:
	//   200: GetCarsNumResponse
	router.Path("/parser-api/kolesa/car/count").
		Methods("GET").
		Handler(getCarsNum)

	// swagger:route GET /parser-api/kolesa/average_price Cars GetCarsByAvgPriceRequest
	// Get car adds by average price in range min and max
	// responses:
	//   200: GetCarsByAvgPriceResponse
	router.Path("/parser-api/kolesa/car/average_price").
		Methods("GET").
		Handler(getCarsByAvgPrice)

	// swagger:route GET /parser-api/kolesa/price Cars GetCarsByPriceRequest
	// Get car adds by price in range min and max
	// responses:
	//   200: GetCarsByPriceResponse
	router.Path("/parser-api/kolesa/car/price").
		Methods("GET").
		Handler(getCarsByPrice)

	// swagger:route GET /parser-api/kolesa/average_price Price GetAvgPriceRequest
	// Get average price by mark, model, year
	// responses:
	// 200: GetAvgPriceResponse
	router.Path("/parser-api/kolesa/average_price").
		Methods("GET").
		Handler(getAvgPrice)

	return router
}
