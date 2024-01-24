package handlers

import (
	"database/sql"
	"time"
	"log"
	"net/http"
	"strconv"

	"github.com/OnescuAlex-Vlad/rari/models"
	"github.com/OnescuAlex-Vlad/rari/utils"
	"github.com/OnescuAlex-Vlad/rari/view/components"
	"github.com/labstack/echo/v4"
)

type CarHandler struct{}

func (h CarHandler) GetCarByIdHandler(c echo.Context) error {
	db, err := models.CreateConnection()
	if err != nil {
		log.Println("Error connecting to the database: ", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("Error parsing car ID: ", err)
		return c.String(http.StatusBadRequest, "Invalid car ID")
	}

	row := db.QueryRow("SELECT * FROM car WHERE id = $1", id)
	

	var car models.Car
	err = row.Scan(
		&car.Id,
		&car.Brand,
		&car.Model,
		&car.Year,
		&car.Color,
		&car.Price,
		&car.Engine,
		&car.CreatedAt,
		&car.UpdatedAt,
		&car.DeletedAt,
	)
	if err == sql.ErrNoRows {
		return c.String(http.StatusNotFound, "Car not found")
	} else if err != nil {
		log.Println("Error scanning row: ", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return utils.Render(c, components.ShowCarPage(car))

}

func (h CarHandler) GetCarHandler(c echo.Context) error {
	db, err := models.CreateConnection()
	if err != nil {
		log.Println("Error connecting to the database: ", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM car")
	if err != nil {
		log.Println("Error querying the database: ", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		err := rows.Scan(
			&car.Id,
			&car.Brand,
			&car.Model,
			&car.Year,
			&car.Color,
			&car.Price,
			&car.Engine,
			&car.CreatedAt,
			&car.UpdatedAt,
			&car.DeletedAt,
		)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		cars = append(cars, car)
	}

	// Create a new car instance
	// one_car := models.Car{
	// 	Brand:     "Ferrari",
	// 	Model:     "488",
	// 	Year:      2019,
	// 	Color:     "Red",
	// 	Engine:    "v8",
	// 	Price:     "200000",
	// 	CreatedAt: "2021-01-01",
	// 	UpdatedAt: "2021-01-01",
	// 	DeletedAt: "2021-01-01",
	// }

	return utils.Render(c, components.ShowCars(cars))
}

func (h CarHandler) CreateCarHandler(c echo.Context) error {
	db, err := models.CreateConnection()
	if err != nil {
		log.Println("Error connecting to the database: ", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	defer db.Close()

	// Parse car data from the request
	var car models.Car
	if err := c.Bind(&car); err != nil {
		log.Println("Error parsing request body: ", err)
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	// Set CreatedAt timestamp
	car.CreatedAt = time.Now().UTC().Format("2006-01-02T15:04:05")

	// Insert the new car into the database
	result, err := db.Query(`
		INSERT INTO cars (brand, model, year, color, price, engine, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`, car.Brand, car.Model, car.Year, car.Color, car.Price, car.Engine, car.CreatedAt)

	if err != nil {
		log.Println("Error inserting car into the database: ", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Retrieve the generated ID
	var newCarID int
	if err := result.Scan(&newCarID); err != nil {
		log.Println("Error retrieving generated ID: ", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Set the generated ID in the response
	car.Id = newCarID

	return c.JSON(http.StatusCreated, car)

}