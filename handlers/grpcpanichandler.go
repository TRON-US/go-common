package handlers

import (
panichandler "github.com/kazegusuri/grpc-panic-handler"
"google.golang.org/grpc"
)

var (
	UIntOpt = grpc.UnaryInterceptor(panichandler.UnaryPanicHandler)
	SIntOpt = grpc.StreamInterceptor(panichandler.StreamPanicHandler)
)

