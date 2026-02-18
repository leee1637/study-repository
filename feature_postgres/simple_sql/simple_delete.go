package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQery := `
	DELETE FROM tasks
	WHERE id = 7;	
	`

	_, err := conn.Exec(ctx, sqlQery)

	return err
}
