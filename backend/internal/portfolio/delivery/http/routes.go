package http

import (
	"github.com/labstack/echo/v4"

	"github.com/seregaa020292/capitalhub/infrastructure/middleware"
	"github.com/seregaa020292/capitalhub/internal/portfolio"
)

// Map portfolio routes
func MapPortfolioRoutes(portfolioGroup *echo.Group, handler portfolio.Handlers, mw *middleware.MiddlewareManager) {
	portfolioGroup.GET("/active-total", handler.GetActiveTotal(), mw.AuthJWTMiddleware, mw.CSRF)
	portfolioGroup.GET("/all-stats", handler.GetAllStats(), mw.AuthJWTMiddleware, mw.CSRF)
	portfolioGroup.POST("/add", handler.Add(), mw.AuthJWTMiddleware, mw.CSRF)
	portfolioGroup.PUT("/:portfolio_id/choose", handler.Choose(), mw.AuthJWTMiddleware, mw.CSRF)
	portfolioGroup.PUT("/:portfolio_id", handler.Edit(), mw.AuthJWTMiddleware, mw.CSRF)
	portfolioGroup.DELETE("/:portfolio_id", handler.Remove(), mw.AuthJWTMiddleware, mw.CSRF)
}
