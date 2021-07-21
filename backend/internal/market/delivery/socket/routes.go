package socket

import (
	"github.com/labstack/echo/v4"

	"github.com/seregaa020292/capitalhub/infrastructure/middleware"
	"github.com/seregaa020292/capitalhub/internal/market"
)

func MapMarketRoutes(marketGroup *echo.Group, handler market.SocketHandlers, mw *middleware.MiddlewareManager) {
	marketGroup.GET("/quotes", handler.Quotes(), mw.AuthJWTMiddleware)
}
