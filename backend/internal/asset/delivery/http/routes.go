package http

import (
	"github.com/labstack/echo/v4"

	"github.com/seregaa020292/capitalhub/infrastructure/middleware"
	"github.com/seregaa020292/capitalhub/internal/asset"
)

// Map asset routes
func MapAssetRoutes(assetGroup *echo.Group, h asset.Handlers, mw *middleware.MiddlewareManager) {
	assetGroup.POST("/add", h.Add(), mw.AuthJWTMiddleware, mw.CSRF)
	assetGroup.GET("/all", h.GetAll(), mw.AuthJWTMiddleware, mw.CSRF)
	assetGroup.GET("/total-all", h.GetTotalAll(), mw.AuthJWTMiddleware, mw.CSRF)
	assetGroup.DELETE("/:asset_id", h.Delete(), mw.AuthJWTMiddleware, mw.CSRF)
	assetGroup.PUT("/:asset_id", h.Update(), mw.AuthJWTMiddleware, mw.CSRF)
	assetGroup.GET("/:asset_id", h.GetByID(), mw.AuthJWTMiddleware, mw.CSRF)
	assetGroup.GET("/market/:market_id", h.GetAllByMarketID(), mw.AuthJWTMiddleware, mw.CSRF)
}
