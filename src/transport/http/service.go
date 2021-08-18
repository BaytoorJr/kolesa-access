package http

import (
	"net/http"

	httpencoders "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/transport/http"
	"github.com/go-kit/kit/log"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/BaytoorJr/kolesa-access/src/middleware"
)

// HTTP service constructor
func NewHTTPService(svcEndpoints *middleware.Endpoints, options []kithttp.ServerOption, logger log.Logger) http.Handler {
	errorEncoder := kithttp.ServerErrorEncoder(
		httpencoders.EncodeErrorResponse,
	)

	errorLogger := kithttp.ServerErrorHandler(
		kittransport.NewLogErrorHandler(logger),
	)

	options = append(options, errorEncoder, errorLogger)

	return initializeRoutes(svcEndpoints, options)
}
