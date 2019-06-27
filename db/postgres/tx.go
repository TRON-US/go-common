package postgres

import (
	"context"

	"github.com/tron-us/go-common/common"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// RunInTransactionContext wraps around underlying go-pg's rollback-supported transaction execution
// with our custom context so it can be easily passed down.
func (db *TGPGDB) RunInTransactionContext(ctx context.Context, txFunc func(context.Context) error) error {
	return db.DB.RunInTransaction(func(tx *pg.Tx) error {
		// Pass ctx with tx object down to the transaction execution
		return txFunc(context.WithValue(ctx, common.ContextPostgresTx, tx))
	})
}

// The following functions are necessary to override include to support both transaction
// and transaction-less queries through the ctx's tx existence.

func (db *TGPGDB) ModelContext(ctx context.Context, models ...interface{}) *orm.Query {
	if tx, ok := ctx.Value(common.ContextPostgresTx).(*pg.Tx); ok {
		return tx.ModelContext(ctx, models...)
	} else {
		return db.DB.ModelContext(ctx, models...)
	}
}

func (db *TGPGDB) ExecContext(ctx context.Context, query interface{}, params ...interface{}) (pg.Result, error) {
	if tx, ok := ctx.Value(common.ContextPostgresTx).(*pg.Tx); ok {
		return tx.ExecContext(ctx, query, params...)
	} else {
		return db.DB.ExecContext(ctx, query, params...)
	}
}
