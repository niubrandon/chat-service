package server

import (
	"context"

	pb "github.com/niubrandon/chat-service/proto/proto/chatgpt/v1"
)

// Server implements the protobuf interface
type Server struct {
	pb.UnimplementedChatGPTServer
}

// New initializes a new Server struct.
func New() *Server {
	return &Server{}
}

// adds a user to the in-memory store.
func (b *Server) GenerateResponse(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	// Call ChatGPT with in.Input
	// Replace this with actual ChatGPT interaction logic
	response := "Placeholder response from ChatGPT"
	return &pb.Response{Output: response}, nil
}
