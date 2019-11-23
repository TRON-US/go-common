package middleware

import (
	"bytes"
	"errors"
	"google.golang.org/grpc"
	"runtime/debug"

	"github.com/tron-us/go-common/v2/log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// Panic handler prints the stack trace when recovering from a panic.
	RecoveryCustomFunc grpc_recovery.RecoveryHandlerFunc = grpc_recovery.RecoveryHandlerFunc(func(p interface{}) error {
		buf := bytes.NewBuffer(debug.Stack())
		log.Error("Panic attack :", zap.Error(errors.New(buf.String())))
		return status.Errorf(codes.Internal, "%s", p)
	})
	// Shared options for the logger, with a custom gRPC code to log level function.
	Opts = []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(RecoveryCustomFunc),
	}
	UnaryServerInterceptor = grpc_recovery.UnaryServerInterceptor(Opts...)
	GrpcServerOption       grpc.ServerOption
)

func init() {
	GrpcServerOption = grpc_middleware.WithUnaryServerChain(UnaryServerInterceptor)
}
