package skelmi

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ErrorBadRequest() *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
}

func ErrorUnauthorized() *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
}

func ErrorForbidden() *echo.HTTPError {
	return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
}

func ErrorNotFound() *echo.HTTPError {
	return echo.NewHTTPError(http.StatusNotFound, "Not found")
}

func ErrorConflict() *echo.HTTPError {
	return echo.NewHTTPError(http.StatusConflict, "Conflict")
}

func ErrorUnprocessableEntity() *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, "Unprocessable entity")
}

func ErrorInternal(message string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, message)
}
