package skelmi

import "github.com/labstack/echo/v4"

// Block represents a logical block, e.g. session, user etc.
type Block interface {
	GetMiddlewares() []echo.MiddlewareFunc
	GetHandlers() []Handler
}
