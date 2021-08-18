module github.com/BaytoorJr/kolesa-access

go 1.5

require (
	git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib v0.0.1
	github.com/go-kit/kit v0.10.0
	github.com/gorilla/mux v1.7.3
	github.com/jackc/pgx/v4 v4.9.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/prometheus/client_golang v1.3.0
)

replace git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib => git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib.git v0.0.1
