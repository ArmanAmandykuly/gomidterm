package task

import (
	// "fmt"
	"log"
	"github.com/ArmanAmandykuly/gomidterm/pkg/database/postgres"
)

func GetTasks() ([]Task, error) {
	rows, err := postgres.DB.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Title, &t.Content)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func GetTaskById(id int) (Task, error) {
	query := "SELECT * FROM tasks WHERE id = $1"
    rows, err := postgres.DB.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if rows.Next() {
        var t Task
        if err := rows.Scan(&t.ID, &t.Title, &t.Content); err != nil {
            log.Println("Error scanning row:", err)
            return Task{}, err
        }
        return t, nil
    }

	var t Task

	return t, nil
}

func SaveTask(task Task) (Task, error) {
	query := "INSERT INTO tasks VALUES($1, $2, $3)"
	_, err := postgres.DB.Query(query, task.ID, task.Title, task.Content)

	if(err != nil) {
		log.Fatal("Error happened while trying to post the task")
		return task, err
	}

	return task, nil
}