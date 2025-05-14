package server

import (
	"context"
	"fmt"
	"net"
	"yim-go/api/internal/core/service"
	"yim-go/api/internal/handler"
	"yim-go/api/internal/repositories/repository"
	"yim-go/api/property"
	"yim-go/shared/gen/protobuf"
	"yim-go/shared/infrastructure"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() error {
	ctx := ContextWithSignal(context.Background())

	secretClient := infrastructure.NewSecretManagerClient(ctx)
	property.Init(ctx, secretClient)

	// init infrastructure
	var (
		pgx, scanapi = infrastructure.NewPostgres(ctx, property.Get().Postgres.PostgresConfig)
		sqlbuilderFlavor = infrastructure.NewQueryBuilder()
	)

	// init repository
	var (
		repo = repository.New(pgx, scanapi, sqlbuilderFlavor)
	)

	// init service
	svc := service.New(repo)

	// create the gRPc server
	s := grpc.NewServer()

	// init service server
	serverServer := handler.New(svc)

	// register gRPC server handler
	protobuf.RegisterBookServiceServer(s, serverServer)

	addr := fmt.Sprintf(":%s", property.Get().Server.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// Use a goroutine to handle the server shutdown
	errChan := make(chan error)
	go func() {
		defer s.GracefulStop()
		fmt.Printf("Server listening on port %s\n", addr)
		if err := s.Serve(lis); err != nil {
			errChan <- err
		}
	}()

	// Register reflection service on gRPC server
	if property.Get().Server.ServerReflection {
		reflection.Register(s)
	}

	// ==================== Signal block ====================
	// block until either context, OS signal, or server fatal error
	select {
	case err := <-errChan:
		fmt.Printf("Error at server: %v\n", err)
		break
	case <-ctx.Done():
		fmt.Println("Stopped")
		break
	}
	return nil
}