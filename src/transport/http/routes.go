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

	router := mux.NewRouter()

	// swagger:route GET /parser-api/kolesa/mark Cars GetCarsByMarkRequest
	// Get car adds by mark
	// responses:
	//   200: GetCarsByMarkResponse
	router.Path("/parser-api/kolesa/mark").
		Methods("GET").
		Handler(getCarsByMark)

	// swagger:route GET /parser-api/kolesa/mark_model Cars GetCarsByMarkAndModel
	// Get car adds by mark and model
	// responses:
	//   200: GetCarsByMarkAndModelResponse
	router.Path("/parser-api/kolesa/mark&model").
		Methods("GET").
		Handler(getCarsByMarkAndModel)

	return router
}
