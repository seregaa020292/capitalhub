package http

import (
	"github.com/labstack/echo/v4"

	"github.com/seregaa020292/capitalhub/infrastructure/middleware"
	"github.com/seregaa020292/capitalhub/internal/market"
)

// Map market routes
func MapMarketRoutes(marketGroup *echo.Group, h market.Handlers, mw *middleware.MiddlewareManager) {
	marketGroup.POST("/create", h.Create(), mw.AuthJWTMiddleware, mw.CSRF)
	marketGroup.PUT("/:market_id", h.Update(), mw.AuthJWTMiddleware, mw.CSRF)
	marketGroup.DELETE("/:market_id", h.Delete(), mw.AuthJWTMiddleware, mw.CSRF)
	marketGroup.GET("/:market_id", h.GetByID(), mw.AuthJWTMiddleware, mw.CSRF)
	marketGroup.GET("/search", h.SearchByTitle(), mw.AuthJWTMiddleware, mw.CSRF)
	marketGroup.GET("/all", h.GetAll(), mw.AuthJWTMiddleware, mw.CSRF)
	marketGroup.GET("/parse", h.Parse(), mw.AuthJWTMiddleware)
}
