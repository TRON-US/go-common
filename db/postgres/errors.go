package postgres

import (
	"strings"
)

const (
	DBMigrationsConcurrentInitErrPrefix = "ERROR #23505" // safe to ignore, error when gopg_migrations already exists
	DBStatementCancelledErrPrefix       = "ERROR #57014"
)

func DBMigrationsAlreadyInit(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), DBMigrationsConcurrentInitErrPrefix)
}

func DBStatementCancelledByUser(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), DBStatementCancelledErrPrefix)
}
