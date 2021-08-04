package application

import "github.com/labstack/echo/v4"

// Application HTTP Handlers interface
type Handlers interface {
	GetDashboard() echo.HandlerFunc
}
