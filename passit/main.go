package main

import (
    "context"
    "log"
    "time"
    proto "github.com/ncostamagna/passit-proto/go/grpcPassit"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func main() {
    // Create a connection to the gRPC server
    conn, err := grpc.Dial("localhost:8050", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Create a client
    client := proto.NewPassitClient(conn)

    // Create a context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // create secret
	createSecret(ctx, client)
	
}

func createSecret(ctx context.Context, client proto.PassitClient) {
    request := &proto.CreateSecretRequest{
        Message: "secret encrypted",
        OneTime: true,
        Expiration: 3600,
    }

    // Make the request
    response, err := client.CreateSecret(ctx, request)
    if err != nil {
        log.Fatalf("Failed to create payment intent: %v", err)
    }

    // Print the response
    log.Printf("Response: %v", response)
}