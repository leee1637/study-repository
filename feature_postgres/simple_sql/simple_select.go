package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at
	FROM tasks
	WHERE completed = FALSE
	ORDER BY id ASC
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel

		rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Created_at,
			&task.Completed_at,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)

		PrintTask(task)
	}

	fmt.Println(tasks[1])

	return tasks, nil
}

func PrintTask(task TaskModel) {
	fmt.Println("-----------")
	fmt.Println("ID: ", task.Id)
	fmt.Println("title: ", task.Title)
	fmt.Println("Description", task.Description)
	fmt.Println("Completed", task.Completed)
}
