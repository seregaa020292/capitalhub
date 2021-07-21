package portfolio

import "github.com/labstack/echo/v4"

// Portfolio HTTP Handlers interface
type Handlers interface {
	GetActiveTotal() echo.HandlerFunc
}
