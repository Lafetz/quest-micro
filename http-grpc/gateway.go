package httpgrpc

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	knightpb "github.com/lafetz/quest-micro/proto/gen"
	"google.golang.org/grpc"
)

func withLogger(logger *slog.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Run request", "http_method", r.Method, "http_url", r.URL)

		h.ServeHTTP(w, r)
	})
}

func NewGatewayServer(conn *grpc.ClientConn, port int, logger *slog.Logger) (*http.Server, error) {
	gwmux := runtime.NewServeMux(

		runtime.WithErrorHandler(HttpErrorHandler),
	)
	err := knightpb.RegisterKnightServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: withLogger(logger, gwmux),
	}, nil
}
