package main

import (
	"github.com/OnescuAlex-Vlad/rari/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	carHandler := handlers.CarHandler{}
	e.GET("/api/v1/cars", carHandler.GetCarHandler)
	e.GET("/api/v1/car/:id", carHandler.GetCarByIdHandler)
	e.POST("/api/v1/car", carHandler.CreateCarHandler)

	e.Logger.Fatal(e.Start(":4200"))
}
