package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	commongrpc "github.com/lafetz/quest-micro/common/grpc"
	httpgrpc "github.com/lafetz/quest-micro/http-grpc"
)

func main() {
	conn, err := commongrpc.NewGRPCClient("localhost:8080")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	srv, err := httpgrpc.NewGatewayServer(conn, 3000, slog.Default())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	if err := srv.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to listen and serve", err)
	}

}
