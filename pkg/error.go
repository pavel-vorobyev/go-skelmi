package skelmi

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ErrorBadRequest(message string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, message)
}

func ErrorUnauthorized(message string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnauthorized, message)
}

func ErrorForbidden(message string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusForbidden, message)
}

func ErrorNotFound(message string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusNotFound, message)
}

func ErrorConflict(message string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusConflict, message)
}

func ErrorUnprocessableEntity(message string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, message)
}

func ErrorInternal(message string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, message)
}
