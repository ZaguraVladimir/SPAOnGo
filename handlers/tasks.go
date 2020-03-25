package handlers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"SPAOnGo/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	data, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//fmt.Fprintf(w, "GET, %q", html.EscapeString(r.URL.Path))
	} else if r.Method == "PUT" {
		//fmt.Fprintf(w, "POST, %q", html.EscapeString(r.URL.Path))
	} else {
		http.Error(w, "Invalid request method.", 405)
	}

	/* 	w.Header().Set("Content-Type", "text/html")
	   	w.WriteHeader(http.StatusOK)
	   	data, err:  = ioutil.ReadFile("public/index.html")
	   	if err ! = nil {
	   		panic(err)
	   	}
	   	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	   	fmt.Fprint(w, string(data) )
	*/
}

//--------------------------------------------------------------------------------------------
// конечная точка GetTasks
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// получаем задачи из модели
		return c.JSON(http.StatusOK, models.GetTasks())
	}
}

// конечная точка PutTask
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// создаём новую задачу
		var task models.Task
		// привязываем пришедший JSON в новую задачу
		c.Bind(&task)
		// добавим задачу с помощью модели
		id, err := models.PutTask(task.Name)
		// вернём ответ JSON при успехе
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// обработка ошибок
		} else {
			return err
		}
	}
}

/// конечная точка DeleteTask
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// используем модель для удаления задачи
		_, err := models.DeleteTask(id)
		// вернём ответ JSON при успехе
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// обработка ошибок
		} else {
			return err
		}
	}
}
