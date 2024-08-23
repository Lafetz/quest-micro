package httpgrpc

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	knightv1 "github.com/lafetz/quest-micro/proto/gen/knight/v1"

	"google.golang.org/grpc"
)

func recoverPanic(logger *slog.Logger, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {

			if err := recover(); err != nil {
				var errorMessage string
				if e, ok := err.(error); ok {
					errorMessage = e.Error()
				} else {
					errorMessage = fmt.Sprintf("unknown panic: %v", err)
				}
				logger.Error("recovered from panic",
					"method", r.Method,
					"url", r.URL.String(),
					"error", errorMessage,
					"stack_trace", string(debug.Stack()),
				)
			}
		}()
		next.ServeHTTP(w, r)
	}
}

func NewGatewayServer(conn *grpc.ClientConn, port int, logger *slog.Logger) (*http.Server, error) {
	gwmux := runtime.NewServeMux(

		runtime.WithErrorHandler(HttpErrorHandler),
	)

	err := knightv1.RegisterKnightServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: recoverPanic(logger, gwmux),
	}, nil
}
