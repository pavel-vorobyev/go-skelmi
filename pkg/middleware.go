package skelmi

import "github.com/labstack/echo/v4"

type Middleware interface {
	Relay(next echo.HandlerFunc) echo.HandlerFunc
}
