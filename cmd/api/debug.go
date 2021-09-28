package main

import (
	"github.com/brotherhood228/dating-bot-api/internal/health"
	"github.com/brotherhood228/dating-bot-api/pkg/metric"
	"github.com/labstack/echo"
)

//StartDebugService run debug port
func StartDebugService(port string) {
	server := echo.New()

	server.GET("/metrics", echo.WrapHandler(metric.Handler()))

	server.GET("/health", health.Check)

	server.Start(port)
}
