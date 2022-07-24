package psql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

type PsqlConfig struct {
	Username, Password, Host, Port, Database string
	MaxConnRetry                             int
	RetryTimeout                             time.Duration
	RetNeedSleep                             time.Duration
}

func NewPsqlClient(ctx context.Context, cfg *PsqlConfig) (connPool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	fmt.Println(dsn)
	retry := cfg.MaxConnRetry

	for retry > 0 {
		ctx, cancel := context.WithTimeout(ctx, cfg.RetryTimeout)
		defer cancel()

		connPool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			time.Sleep(cfg.RetNeedSleep)
			retry--
			continue
		}
		return
	}
	return
}
