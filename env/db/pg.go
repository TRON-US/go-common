package db

import (
	"fmt"
	"os/user"
	"strconv"
	"time"

	"github.com/tron-us/go-common/common"
	"github.com/tron-us/go-common/env"
	"github.com/tron-us/go-common/log"

	"go.uber.org/zap"
)

const (
	DBReadTimeout  = 1 * time.Minute
	DBWriteTimeout = 1 * time.Minute
)

var (
	DBReadURL       string // https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
	DBWriteURL      string
	DBMigrationsDir = "./migrations"
	DBStmtTimeout   = 10 * time.Second
	DBNumConns      = 0 // max db conns, 0 = use driver-recommended default
)

func init() {
	// Fetch master db as default write, can't be empty;
	envKey, duW := env.GetEnv("DB_URL")
	if duW != "" {
		DBWriteURL = duW
	} else if u, err := user.Current(); err == nil {
		// Attempt at connecting to a common local database
		DBWriteURL = fmt.Sprintf("postgresql://%s@localhost:5432/%s", u.Username, u.Username)
	} else {
		log.Panic(common.EmptyVarError, zap.String("env", envKey))
	}
	// if slave url passed, use it as read default
	if _, duR := env.GetEnv("DB_READ_URL"); duR != "" {
		DBReadURL = duR
	} else {
		DBReadURL = DBWriteURL
	}
	if _, md := env.GetEnv("MIGRATIONS_DIR"); md != "" {
		DBMigrationsDir = md
	}
	if envKey, dbst := env.GetEnv("DB_STMT_TIMEOUT"); dbst != "" {
		if toInt, err := strconv.ParseInt(dbst, 10, 64); err != nil {
			log.Warn(common.IntConversionError, zap.String("env", envKey), zap.Error(err))
		} else {
			DBStmtTimeout = time.Duration(toInt) * time.Second
		}
	}
	if envKey, dbnc := env.GetEnv("DB_NUM_CONNS"); dbnc != "" {
		if toInt, err := strconv.ParseInt(dbnc, 10, 64); err != nil {
			log.Warn(common.IntConversionError, zap.String("env", envKey), zap.Error(err))
		} else {
			DBNumConns = int(toInt)
		}
	}
}
