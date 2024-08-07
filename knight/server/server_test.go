package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/lafetz/quest-micro/common/logger"
	knight "github.com/lafetz/quest-micro/knight/core"
	"github.com/lafetz/quest-micro/knight/repository"
	protoknight "github.com/lafetz/quest-micro/proto/knight"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

var mongoC *mongo.Client
var cleanupFunc func()

func TestMain(m *testing.M) {
	ctx := context.Background()
	mongodbContainer, err := mongodb.Run(ctx, "mongo:7.0.5")
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	ip, err := mongodbContainer.Host(context.TODO())
	if err != nil {
		log.Fatalf("failed to get container host: %s", err)
	}

	mappedPort, err := mongodbContainer.MappedPort(context.TODO(), "27017")
	if err != nil {
		log.Fatalf("failed to get mapped port: %s", err)
	}
	uri := fmt.Sprintf("mongodb://%s:%s", ip, mappedPort.Port())
	logger := logger.NewLogger("dev")
	mongo, close, err := repository.NewDb(uri, logger)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %s", err)
	}
	mongoC = mongo
	cleanupFunc = func() {
		close()
		if err := mongodbContainer.Terminate(context.TODO()); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}

	code := m.Run()

	cleanupFunc()
	os.Exit(code)
}
func newServer(t *testing.T, register func(srv *grpc.Server)) *grpc.ClientConn {
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop()
	})

	register(srv)

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("srv.Serve %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	t.Cleanup(func() {
		conn.Close()
	})
	if err != nil {
		t.Fatalf("grpc.DialContext %v", err)
	}

	return conn
}
func TestAddKnight(t *testing.T) {
	store, err := repository.NewStore(mongoC)
	if err != nil {
		t.Fatalf("failed to create store: %s", err)
	}
	svc := knight.NewKnightService(store)
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	grpcServer := NewGrpcServer(svc, 8080, logger.NewLogger("dev"))

	conn := newServer(t, func(srv *grpc.Server) {
		protoknight.RegisterKnightServiceServer(srv, grpcServer)
	})

	client := protoknight.NewKnightServiceClient(conn)
	type expectation struct {
		out *protoknight.AddKnightRes
		err error
	}
	tests := map[string]struct {
		in       *protoknight.AddKnightReq
		expected expectation
	}{
		"Successful AddKnight": {
			in: &protoknight.AddKnightReq{
				Name:  "Sir Lancelot",
				Email: "lancelot@example.com",
			},
			expected: expectation{
				out: &protoknight.AddKnightRes{
					Id:       "some-uuid", // UUID validation is expected
					Name:     "Sir Lancelot",
					Email:    "lancelot@example.com",
					IsActive: true,
				},
				err: nil,
			},
		},
		"Invalid Email": {
			in: &protoknight.AddKnightReq{
				Name:  "Sir Lancelot",
				Email: "invalid-email",
			},
			expected: expectation{
				out: nil,
				err: status.Error(codes.InvalidArgument, "invalid knight request"),
			},
		},
		"Empty Name": {
			in: &protoknight.AddKnightReq{
				Name:  "",
				Email: "lancelot@example.com",
			},
			expected: expectation{
				out: nil,
				err: status.Error(codes.InvalidArgument, "invalid knight request"),
			},
		},
		"Empty Email": {
			in: &protoknight.AddKnightReq{
				Name:  "Sir Lancelot",
				Email: "",
			},
			expected: expectation{
				out: nil,
				err: status.Error(codes.InvalidArgument, "invalid knight request"),
			},
		},
		"Duplicate Email": {
			in: &protoknight.AddKnightReq{
				Name:  "Sir Lancelot",
				Email: "lancelot@example.com",
			},
			expected: expectation{
				out: nil,
				err: status.Error(codes.AlreadyExists, knight.ErrEmailUnique.Error()),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			res, err := client.AddKnight(context.Background(), tt.in)

			if err != nil {
				if tt.expected.err == nil || err.Error() != tt.expected.err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.err != nil {
					t.Errorf("Expected error: %v, but got none", tt.expected.err)
				}
				if res == nil {
					t.Fatalf("Expected non-nil response, got nil")
				}
				if tt.expected.out != nil {
					if res.Name != tt.expected.out.Name ||
						res.Email != tt.expected.out.Email ||
						res.IsActive != tt.expected.out.IsActive {
						t.Errorf("Unexpected response: got %+v, want %+v", res, tt.expected.out)
					}
				}

			}
		})
	}
}
