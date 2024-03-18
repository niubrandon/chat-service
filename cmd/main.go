package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	chatgptv1 "github.com/niubrandon/chat-service/proto/proto/chatgpt/v1"
	"github.com/niubrandon/chat-service/server"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint

	server := server.New()

	grpcServer := grpc.NewServer()
	chatgptv1.RegisterChatGPTServer(grpcServer, server)

	reflection.Register(grpcServer)
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listen.Close()
	go func() {
		log.Println("gRPC server started on port :50051")
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = chatgptv1.RegisterChatGPTHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)

	if err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
