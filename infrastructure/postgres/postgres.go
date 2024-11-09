package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/kbiits/dealls-take-home-test/config"
)

func ConnectToPostgres(
	dbConfig config.DatabaseConfig,
) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Name,
	)

	connConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	// test on borrow
	afterConnect := stdlib.OptionAfterConnect(func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(ctx, `SELECT 1`)
		if err != nil {
			return err
		}
		return nil
	})

	pgxdb := stdlib.OpenDB(*connConfig, afterConnect)
	return sqlx.NewDb(pgxdb, "pgx"), nil
}
