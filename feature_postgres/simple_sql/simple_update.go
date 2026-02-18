package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQery := `
	UPDATE tasks
	SET description = ')))'
	WHERE completed = FALSE;
	`

	_, err := conn.Exec(ctx, sqlQery)

	return err
}
