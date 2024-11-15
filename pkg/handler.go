package skelmi

import "github.com/labstack/echo/v4"

// Handler represents a handler function to be registered in Echo instance
type Handler struct {
	Method      string
	Path        string
	Middlewares []echo.MiddlewareFunc
	Handle      echo.HandlerFunc
}
