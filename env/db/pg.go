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
	DBNumConns      = 0 // max db conns, 0 = use driver-recommended default\
	readUserName    string
	readPwd         string
	readHost        string
	readDbName      string
	writeUserName   string
	writePwd        string
	writeHost       string
	writeDbName     string
	un              string
	pwd             string
	host            string
	dbName          string
)

func init() {
	// Fetch master db as default write, can't be empty;
	_, duW := env.GetEnv("DB_URL")
	if duW != "" {
		DBWriteURL = duW
	} else {
		envDbWUnKey, _un := env.GetEnv("DB_WRITE_USERNAME")
		if _un != "" {
			writeUserName = _un
		} else {
			log.Panic(constant.EmptyVarError, zap.String("env", envDbWUnKey))
		}

		envDbWPwdKey, _pwd := env.GetEnv("DB_WRITE_PASSWORD")
		if _pwd != "" {
			writePwd = _pwd
		} else {
			log.Panic(constant.EmptyVarError, zap.String("env", envDbWPwdKey))
		}

		envDbWHost, _host := env.GetEnv("DB_WRITE_HOST")
		if _host != "" {
			writeHost = _host
		} else {
			log.Panic(constant.EmptyVarError, zap.String("env", envDbWHost))
		}

		envDbWNameKey, _dbName := env.GetEnv("DB_WRITE_NAME")
		if _dbName != "" {
			writeDbName = _dbName
		} else {
			log.Panic(constant.EmptyVarError, zap.String("env", envDbWNameKey))
		}

	}

	if _, duR := env.GetEnv("DB_READ_URL"); duR != "" {
		DBReadURL = duR
	} else {
		envDbRUnKey, un := env.GetEnv("DB_READ_USERNAME")
		if un != "" {
			readUserName = un
		} else {
			log.Panic(constant.EmptyVarError, zap.String("env", envDbRUnKey))
		}

		envDbRPwdKey, pwd := env.GetEnv("DB_READ_PASSWORD")
		if pwd != "" {
			readPwd = pwd
		} else {
			log.Panic(constant.EmptyVarError, zap.String("env", envDbRPwdKey))
		}

		envDbRHostKey, host := env.GetEnv("DB_READ_HOST")
		if host != "" {
			readHost = host
		} else {
			log.Panic(constant.EmptyVarError, zap.String("env", envDbRHostKey))
		}

		envDbRNameKey, dbName := env.GetEnv("DB_READ_NAME")
		if dbName != "" {
			readDbName = dbName
		} else {
			log.Panic(constant.EmptyVarError, zap.String("env", envDbRNameKey))
		}

	}

	// if slave url passed, use it as read default
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

	DBWriteURL = "postgresql://" + writeUserName + ":" + writePwd + "@" + writeHost + ":5432/" + writeDbName
	DBReadURL = "postgresql://" + readUserName + ":" + readPwd + "@" + readHost + ":5432/" + readDbName

}
