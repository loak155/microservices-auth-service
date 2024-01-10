package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/loak155/microservices-auth-service/client"
	"github.com/loak155/microservices-auth-service/router"
	"github.com/loak155/microservices-auth-service/usecase"
	"github.com/loak155/microservices-auth-service/utils"
	"google.golang.org/grpc"
)

func main() {
	slog.Info("starting grpc server")

	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	server := grpc.NewServer()
	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	userGRPCClient, err := client.NewUserGRPCClient()
	if err != nil {
		slog.Error("failed to connect to user service")
		cancel()
	}
	jwtManager := utils.NewJwtManager(os.Getenv("SECRET_KEY"), time.Hour*12)
	authUsecase := usecase.NewAuthUsecase(userGRPCClient, *jwtManager)
	router.NewAuthGRPCServer(server, authUsecase)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
	if err != nil {
		slog.Error("failed to listen to address")
		cancel()
	}
	err = server.Serve(listener)
	if err != nil {
		slog.Error("failed to start gRPC server")
		cancel()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case v := <-quit:
		slog.Info("signal.Notify: ", v)
	case done := <-ctx.Done():
		slog.Info("ctx.Done: ", done)
	}
}
