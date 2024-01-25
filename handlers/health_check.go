package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func HealthCheck(c echo.Context) error {
	resp := HealthCheckResponse{
		Message: "Everything is ok!",
	}
	return c.JSON(http.StatusOK, resp)
}