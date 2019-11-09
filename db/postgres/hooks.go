package postgres

import (
	"context"

	"github.com/tron-us/go-common/log"

	"github.com/go-pg/pg/v9"
	"go.uber.org/zap"
)

type dbQueryLoggerHook struct {
	beforeEnabled bool
	afterEnabled  bool
}

func (d dbQueryLoggerHook) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	if !d.beforeEnabled {
		return ctx, nil
	}
	qs, err := q.FormattedQuery()
	if err != nil {
		log.HandlerDebug(ctx, "BeforeQueryLog: Unable to log raw query")
		return ctx, err
	} else {
		log.HandlerDebug(ctx, "BeforeQueryLog: Raw query", zap.String("q", qs))
		return ctx, nil
	}
}

func (d dbQueryLoggerHook) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	if !d.afterEnabled {
		return nil
	}
	qs, err := q.FormattedQuery()
	if err != nil {
		log.HandlerDebug(ctx, "AfterQueryLog: Unable to log raw query")
		return err
	} else {
		log.HandlerDebug(ctx, "AfterQueryLog: Raw query", zap.String("q", qs))
		return nil
	}
}
