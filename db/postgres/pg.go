package postgres

import (
	"context"
	"time"

	"github.com/tron-us/go-common/constant"
	env "github.com/tron-us/go-common/env/db"
	"github.com/tron-us/go-common/log"

	"github.com/go-pg/migrations/v7"
	"github.com/go-pg/pg/v9"
	"go.uber.org/zap"
)

type TGPGDB struct {
	*pg.DB
}

func NewTGPGDB(db *pg.DB) *TGPGDB {
	return &TGPGDB{db}
}

func CreateTGPGDB(url string) *TGPGDB {
	opts, err := pg.ParseURL(url)
	if err != nil {
		log.Panic(constant.DBURLParseError, zap.String("URL:", url), zap.Error(err))
	}
	opts.ReadTimeout = env.DBReadTimeout
	opts.WriteTimeout = env.DBWriteTimeout
	opts.TLSConfig = nil // disabled for faster local connection (even in production)
	if env.DBNumConns > 0 {
		opts.PoolSize = env.DBNumConns
	}

	return NewTGPGDB(pg.Connect(opts))
}

// Ping simulates a "blank query" behavior similar to lib/pq's
// to check if the db connection is alive.
func (db *TGPGDB) Ping() error {
	_, err := db.ExecOne("SELECT 1")
	return err
}

// Migrate check and migrate to lastest db version.
func (db *TGPGDB) Migrate() error {
	// Make sure to only search specified migrations dir
	cl := migrations.NewCollection()
	cl.DisableSQLAutodiscover(true)
	err := cl.DiscoverSQLMigrations(env.DBMigrationsDir)
	if err != nil {
		return err
	}

	// Intentionally ignore harmless errors on initializing gopg_migrations
	_, _, err = cl.Run(db, "init")
	if err != nil && !DBMigrationsAlreadyInit(err) {
		return err
	}

	var oldVersion, newVersion int64
	// Run all migrations in a transaction so we rollback if migrations fail anywhere
	err = db.RunInTransaction(func(tx *pg.Tx) error {
		oldVersion, newVersion, err = cl.Run(db, "up")
		return err
	})
	if err != nil {
		return err
	}
	if newVersion == oldVersion {
		log.Info("db schema up to date")
	} else {
		log.Info("db schema migrated successfully", zap.Int64("from", oldVersion), zap.Int64("to", newVersion))
	}
	return nil
}

// WithContextTimeout executes statements with a default timeout
func WithContextTimeout(ctx context.Context, f func(ctx context.Context)) {
	WithContextTimeoutValue(ctx, env.DBStmtTimeout, f)
}

// WithContextTimeoutValue executes an inner routine while passing a ctx that supports custom
// timeout and query cancellation to the postgres server.
func WithContextTimeoutValue(ctx context.Context, timeout time.Duration, f func(ctx context.Context)) {
	newCtx, cancel := context.WithTimeout(ctx, timeout)
	f(newCtx)
	cancel()
}
