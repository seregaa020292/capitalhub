package asset

import "github.com/labstack/echo/v4"

// Asset HTTP Handlers interface
type Handlers interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetTotalAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	GetAllByMarketID() echo.HandlerFunc
}
