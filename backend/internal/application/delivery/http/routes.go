package http

import (
	"github.com/labstack/echo/v4"

	"github.com/seregaa020292/capitalhub/infrastructure/middleware"
	"github.com/seregaa020292/capitalhub/internal/application"
)

func MapApplicationRoutes(portfolioGroup *echo.Group, handler application.Handlers, mw *middleware.MiddlewareManager) {
	portfolioGroup.GET("/dashboard", handler.GetDashboard(), mw.AuthJWTMiddleware, mw.CSRF)
}
