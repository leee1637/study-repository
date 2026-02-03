package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ChekConnection() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:123@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Подключился к бд")
}
