package main

import (
	"SPAOnGo/models"
	"SPAOnGo/handlers"
	"database/sql"
	"log"
	"net/http"

	//"SPAOnGo/handlers"

	//"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	models.InitDB("storage.db")

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/tasks", handlers.TasksHandler)
	log.Println("Listen and serve on :8080...")
	http.ListenAndServe(":8080", nil)

	/* 	e := echo.New()

	   	e.File("/", "public/index.html")
	   	e.GET("/tasks", handlers.GetTasks(db))
	   	e.PUT("/tasks", handlers.PutTask(db))
	   	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	   	e.Logger.Fatal(e.Start(":8000"))
	*/
}