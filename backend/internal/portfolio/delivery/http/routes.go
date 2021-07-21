package http

import (
	"github.com/labstack/echo/v4"

	"github.com/seregaa020292/capitalhub/infrastructure/middleware"
	"github.com/seregaa020292/capitalhub/internal/portfolio"
)

// Map portfolio routes
func MapPortfolioRoutes(portfolioGroup *echo.Group, handler portfolio.Handlers, mw *middleware.MiddlewareManager) {
	portfolioGroup.GET("/active-total", handler.GetActiveTotal(), mw.AuthJWTMiddleware, mw.CSRF)
}
