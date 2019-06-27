package postgres

import (
	"strings"
)

const (
	DBMigrationsInitializedErrPrefix = "ERROR #42P07" // safe to ignore, error when gopg_migrations already exists
	DBStatementCancelledErrPrefix    = "ERROR #57014"
)

func DBMigrationsAlreadyInit(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), DBMigrationsInitializedErrPrefix)
}

func DBStatementCancelledByUser(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), DBStatementCancelledErrPrefix)
}
