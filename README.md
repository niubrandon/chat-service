# chat-service

## release, and tagging a module with semantic versions
```zsh
go mod tidy
go test ./...
git add .
git commit -m "changes for v1.0.0"
git push origin v1.0.0
```

##
- protoc-gen-grpc-gateway
- protoc-gen-openapiv2
- protoc-gen-go
- protoc-gen-go-grpc

```zsh
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Run locally
```zsh


grpcurl -plaintext -d '{"input": "Hello, ChatGPT!"}' localhost:50051 chatgpt.v1.ChatGPT/GenerateResponse

curl -X POST -d '{"input": "Hello, ChatGPT!"}' http://localhost:8080/v1/chatgpt/generate_response


```