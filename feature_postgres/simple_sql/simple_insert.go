package simple_sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context,
	conn *pgx.Conn,
	title string,
	description string,
	completed bool,
	createdAt time.Time,
) error {
	sqlQery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ($1, $2, $3, $4);
	`
	_, err := conn.Exec(ctx, sqlQery, title, description, completed, createdAt)
	if err != nil {
		return err
	}

	return err
}
