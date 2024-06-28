package rest

import (
	"context"
	"log/slog"
	"net/http"

	servV1 "github.com/GrosbergKirr/proto_contracts/gen/go/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunRest(log *slog.Logger) error {
	ctx, _ := context.WithCancel(context.Background())
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := servV1.RegisterGatewayServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Info("Failed to register gateway", "error", err)
		return err
	}

	log.Info("server listening at 8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Info("Failed to start server", "error", err)
		return err
	}
	return nil
}
