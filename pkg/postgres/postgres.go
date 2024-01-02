package postgres

import (
	"database/sql"
	"fmt"
	"github.com/rodrigoprobst/go-plan-management/pkg/configs"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"go.uber.org/zap"
)

var db *bun.DB

func GetDatabase(logger *zap.Logger) *bun.DB {
	if db == nil {
		// Establish a connection to the PostgreSQL database
		sqlDb := sql.OpenDB(pgdriver.NewConnector(
			pgdriver.WithDSN(configs.PostgresCfg.Dsn),
		))
		db = bun.NewDB(sqlDb, pgdialect.New())

		// Add a query hook for logging
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.FromEnv("BUNDEBUG"),
		))

		// Ping the database to test the connection
		err := db.Ping()
		if err != nil {
			logger.Error(fmt.Sprintf("postgres connection error: %s", err))
			panic(err)
		}
	}
	return db
}
