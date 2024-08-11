package commongrpc

import (
	"log/slog"
	"runtime/debug"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func PanicRecoveryGrpc(logger *slog.Logger) func(p interface{}) (err error) {
	return func(p interface{}) (err error) {
		logger.Error("recovered from panic", "panic", p, "stack", debug.Stack())
		return status.Errorf(codes.Internal, "%v", p)
	}
}
