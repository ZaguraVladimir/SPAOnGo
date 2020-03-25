package models

import (
	"database/sql"
)

var db *sql.DB

// Task это структура с данными задачи
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection это список задач
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func InitDB(filepath string) {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	migrate()
}

func migrate() {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
	);
	`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}

// GetTasks from the DB
func GetTasks() TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)
	// выходим, если SQL не сработал по каким-то причинам
	if err != nil {
		panic(err)
	}
	// убедимся, что всё закроется при выходе из программы
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)
		// выход при ошибке
		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

// PutTask into DB
func PutTask(name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"

	// выполним SQL запрос
	stmt, err := db.Prepare(sql)
	// выход при ошибке
	if err != nil {
		panic(err)
	}
	// убедимся, что всё закроется при выходе из программы
	defer stmt.Close()

	// заменим символ '?' в запросе на 'name'
	result, err2 := stmt.Exec(name)
	// выход при ошибке
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeleteTask from DB
func DeleteTask(id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	// выполним SQL запрос
	stmt, err := db.Prepare(sql)
	// выход при ошибке
	if err != nil {
		panic(err)
	}

	// заменим символ '?' в запросе на 'id'
	result, err2 := stmt.Exec(id)
	// выход при ошибке
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
