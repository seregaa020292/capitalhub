package http

import (
	"github.com/labstack/echo/v4"

	"github.com/seregaa020292/capitalhub/infrastructure/middleware"
	"github.com/seregaa020292/capitalhub/internal/auth"
)

// Map auth routes
func MapAuthRoutes(authGroup *echo.Group, h auth.Handlers, mw *middleware.MiddlewareManager) {
	authGroup.POST("/register", h.Register())
	authGroup.POST("/login", h.Login())
	authGroup.POST("/logout", h.Logout())
	authGroup.POST("/refresh", h.RefreshToken())
	authGroup.GET("/confirmed", h.Confirmed())
	authGroup.GET("/find", h.FindByName())
	authGroup.GET("/all", h.GetUsers())
	authGroup.GET("/:user_id", h.GetUserByID())

	authGroup.Use(mw.AuthJWTMiddleware)
	authGroup.GET("/check", h.CheckLogged())
	authGroup.GET("/token", h.GetCSRFToken())
	authGroup.POST("/:user_id/avatar", h.UploadAvatar(), mw.CSRF)
	authGroup.PUT("/:user_id", h.Update(), mw.OwnerOrAdminMiddleware(), mw.CSRF)
	authGroup.DELETE("/:user_id", h.Delete(), mw.CSRF, mw.RoleBasedAuthMiddleware([]string{"admin"}))
}
