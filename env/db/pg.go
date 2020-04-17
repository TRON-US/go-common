package db

import (
	"strconv"
	"time"

	"github.com/tron-us/go-common/v2/constant"
	"github.com/tron-us/go-common/v2/env"
	"github.com/tron-us/go-common/v2/log"

	"go.uber.org/zap"
)

const (
	DBReadTimeout  = 5 * time.Minute
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

	if _, duW := env.GetEnv("DB_URL"); duW != "" {
		DBWriteURL = duW
		// if slave url passed, use it as read default
		if _, duR := env.GetEnv("DB_READ_URL"); duR != "" {
			DBReadURL = duR
		} else {
			DBReadURL = DBWriteURL
		}
	} else {
		// Fetch master db as default write, can't be empty;
		_, writeUserName := env.GetEnv("DB_WRITE_USERNAME")
		if writeUserName == "" {
			writeUserName = "user"
		}

		_, writePwd := env.GetEnv("DB_WRITE_PASSWORD")
		if writePwd == "" {
			writePwd = "pass"
		}

		_, writeHost := env.GetEnv("DB_WRITE_HOST")
		if writeHost == "" {
			writeHost = "localhost"
		}

		_, writeDbName := env.GetEnv("DB_WRITE_NAME")
		if writeDbName == "" {
			writeDbName = "test_db"
		}

		_, readUserName := env.GetEnv("DB_READ_USERNAME")
		if readUserName == "" {
			readUserName = "user"
		}

		_, readPwd := env.GetEnv("DB_READ_PASSWORD")
		if readPwd == "" {
			readPwd = "pass"
		}

		_, readHost := env.GetEnv("DB_READ_HOST")
		if readHost == "" {
			readHost = "localhost"
		}

		_, readDbName := env.GetEnv("DB_READ_NAME")
		if readDbName == "" {
			readDbName = "test_db"
		}

		DBWriteURL = "postgresql://" + writeUserName + ":" + writePwd + "@" + writeHost + ":5432/" + writeDbName
		DBReadURL = "postgresql://" + readUserName + ":" + readPwd + "@" + readHost + ":5432/" + readDbName
	}

	if _, md := env.GetEnv("MIGRATIONS_DIR"); md != "" {
		DBMigrationsDir = md
	}
	if envKey, dbst := env.GetEnv("DB_STMT_TIMEOUT"); dbst != "" {
		if toInt, err := strconv.ParseInt(dbst, 10, 64); err != nil {
			log.Warn(constant.IntConversionError, zap.String("env", envKey), zap.Error(err))
		} else {
			DBStmtTimeout = time.Duration(toInt) * time.Second
		}
	}
	if envKey, dbnc := env.GetEnv("DB_NUM_CONNS"); dbnc != "" {
		if toInt, err := strconv.ParseInt(dbnc, 10, 64); err != nil {
			log.Warn(constant.IntConversionError, zap.String("env", envKey), zap.Error(err))
		} else {
			DBNumConns = int(toInt)
		}
	}
}
