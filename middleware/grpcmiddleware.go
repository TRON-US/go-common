package middleware

import (
	"errors"
	"runtime"

	"github.com/tron-us/go-common/v2/log"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

var (
	// Panic handler prints the stack trace when recovering from a panic.
	RecoveryCustomFunc grpc_recovery.RecoveryHandlerFunc = grpc_recovery.RecoveryHandlerFunc(func(p interface{}) error {
		buf := make([]byte, 1<<16)
		stacklen := runtime.Stack(buf, true)
		log.Error("Panic attack :", zap.Error(errors.New(string(buf[:stacklen]))))
		return status.Errorf(codes.Internal, "%s", p)
	})
	// Shared options for the logger, with a custom gRPC code to log level function.
	Opts = []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(RecoveryCustomFunc),
	}

	UnaryServerInterceptor = grpc_recovery.UnaryServerInterceptor(Opts...)
)
