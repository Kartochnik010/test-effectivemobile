package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"

	_ "github.com/lib/pq"
)

type Repository interface {
	GetConnection() *pgx.Conn
	Close(ctx context.Context) error
}

type repository struct {
	Client *pgx.Conn
}

func InitDB(ctx context.Context, dsn string) (Repository, error) {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	repo := repository{Client: conn}
	if err = repo.Client.Ping(ctx); err != nil {
		return &repo, err
	}

	return &repo, nil
}

func (c repository) GetConnection() *pgx.Conn {
	return c.Client
}
func (c repository) Close(ctx context.Context) error {
	return c.Client.Close(ctx)
}
