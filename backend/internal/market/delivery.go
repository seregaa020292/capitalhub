package market

import "github.com/labstack/echo/v4"

// Market HTTP Handlers interface
type Handlers interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	SearchByTitle() echo.HandlerFunc
	Parse() echo.HandlerFunc
}

// Market Socket Handlers interface
type SocketHandlers interface {
	Quotes() echo.HandlerFunc
}
