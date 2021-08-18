package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log/level"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	postgresqlConn "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/databases/postgresql"
	httpencoders "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/transport/http"
	corsutil "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/utils/cors"
	healthcheckutil "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/utils/healthcheck"
	liblogger "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/utils/logger"

	"github.com/BaytoorJr/kolesa-access/src/config"
	"github.com/BaytoorJr/kolesa-access/src/constructor"
	carDataMiddleware "github.com/BaytoorJr/kolesa-access/src/middleware"
	"github.com/BaytoorJr/kolesa-access/src/repository/postgresql"
	carDataHTTP "github.com/BaytoorJr/kolesa-access/src/transport/http"
)

func main() {
	// main ctx
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// parse flags
	httpPort := flag.String("http.port", ":8080", "HTTP listen address")
	flag.Parse()

	// init structured logger for api
	logger := liblogger.NewServiceLogger("parser-api")
	_ = level.Info(logger).Log("msg", "service started")

	// init configs
	err := config.InitConfigs()
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	// init postgresql connection
	postgresConn, err := postgresqlConn.InitConnect(
		config.MainConfig.PostgreSQLConfig.PostgresHost,
		config.MainConfig.PostgreSQLConfig.PostgresPort,
		config.MainConfig.PostgreSQLConfig.PostgresUser,
		config.MainConfig.PostgreSQLConfig.PostgresPass,
		config.MainConfig.PostgreSQLConfig.PostgresName)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	// init repositories (data access layer)
	mainRepo, err := postgresql.NewStore(postgresConn, logger)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	// init service (business-logic layer)
	carDataSvc := constructor.NewCarDataService(ctx, mainRepo, logger)

	// init endpoints (endpoints layer)
	carDataEndpoints := carDataMiddleware.MakeEndpoints(carDataSvc)

	// init HTTP handler (transport layer)
	serverOptions := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(httpencoders.EncodeErrorResponse),
	}
	carDataHTTPHandler := carDataHTTP.NewHTTPService(carDataEndpoints, serverOptions, logger)

	// add routes, prometheus and health check handlers
	http.Handle("/parser-api/kolesa/", corsutil.CORS(carDataHTTPHandler))
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/check", healthcheckutil.HealthCheck)

	// init errors chan
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, config.DefaultChanCapacity)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// init HTTP server
	go func() {
		_ = level.Info(logger).Log("transport", "HTTP", "port", *httpPort)
		errs <- http.ListenAndServe(*httpPort, nil)
	}()

	defer func() {
		cancel()
		_ = level.Info(logger).Log("msg", "service ended")
	}()

	_ = level.Error(logger).Log("exit", <-errs)
}
