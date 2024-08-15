package commongrpc

import (
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func NewCb(name string) *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    name,
		Timeout: 5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 5
		},
	})
}
func NewGRPCClient(addr string, cb *gobreaker.CircuitBreaker) (*grpc.ClientConn, error) {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(grpc_retry.WithCodes(codes.Unavailable, codes.Aborted),
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)))))

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	opts = append(opts, grpc.WithUnaryInterceptor(CircuitBreaker(cb)))
	opts = append(opts,
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             time.Second,
			PermitWithoutStream: true,
		}))

	conn, err := grpc.NewClient(addr, opts...)

	if err != nil {

		return nil, err
	}

	return conn, nil
}
