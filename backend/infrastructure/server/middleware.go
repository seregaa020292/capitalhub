package server

import (
	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/infrastructure/session"
	"github.com/seregaa020292/capitalhub/internal/auth"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/seregaa020292/capitalhub/docs"
	apiMiddlewares "github.com/seregaa020292/capitalhub/infrastructure/middleware"
	"github.com/seregaa020292/capitalhub/pkg/csrf"
	"github.com/seregaa020292/capitalhub/pkg/metric"
)

func InitMiddleware(
	e *echo.Echo,
	server *Server,
	sessUC session.UseCase,
	authUC auth.UseCase,
) *apiMiddlewares.MiddlewareManager {
	metrics := createMetrics(server.cfg, server.logger)

	mw := apiMiddlewares.NewMiddlewareManager(sessUC, authUC, server.cfg, []string{"*"}, server.logger)

	e.Use(mw.RequestLoggerMiddleware)

	docs.SwaggerInfo.Title = "REST API"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	if server.cfg.Server.SSL {
		e.Pre(middleware.HTTPSRedirect())
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID, csrf.CSRFHeader},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())
	e.Use(mw.MetricsMiddleware(metrics))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))
	if server.cfg.Server.Debug {
		e.Use(mw.DebugMiddleware)
	}

	return mw
}

func createMetrics(cfg *config.Config, logger logger.Logger) metric.Metrics {
	metrics, err := metric.CreateMetrics(cfg.Metrics.URL, cfg.Metrics.ServiceName)
	if err != nil {
		logger.Errorf("CreateMetrics Error: %s", err)
	}
	logger.Info(
		"Metrics available URL: %s, ServiceName: %s",
		cfg.Metrics.URL,
		cfg.Metrics.ServiceName,
	)

	return metrics
}
