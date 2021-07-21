package auth

import "github.com/labstack/echo/v4"

// Auth HTTP Handlers interface
type Handlers interface {
	Register() echo.HandlerFunc
	Confirmed() echo.HandlerFunc
	Login() echo.HandlerFunc
	Logout() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetUserByID() echo.HandlerFunc
	FindByName() echo.HandlerFunc
	GetUsers() echo.HandlerFunc
	UploadAvatar() echo.HandlerFunc
	GetCSRFToken() echo.HandlerFunc
	CheckLogged() echo.HandlerFunc
	RefreshToken() echo.HandlerFunc
}
