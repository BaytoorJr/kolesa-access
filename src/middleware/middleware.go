package middleware

import "github.com/BaytoorJr/kolesa-access/src/service"

// Middleware describes a service middleware
type Middleware func(service service.CarDataService) service.CarDataService
