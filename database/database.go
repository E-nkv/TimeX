package database

import (
	"context"
	"time"

	pgx "github.com/jackc/pgx/v5"
)

func OpenDB() (*pgx.Conn, error) {

	dsn := "postgresql://postgres:admin@localhost:5432/TimeX?sslmode=disable"

	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if err = db.Ping(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
