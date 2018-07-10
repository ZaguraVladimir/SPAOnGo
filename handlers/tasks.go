package handlers

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type H map[string]interface{}

// конечная точка GetTasks
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "tasks")
	}
}

// конечная точка PutTask
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, H{"created": 123})
	}
}

// конечная точка DeleteTask
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, H{"deleted": id})
	}
}
