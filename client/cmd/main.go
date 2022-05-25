package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"test_grpc/generated/protos/test"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8008", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := test.NewTestServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Hello(ctx, &test.Settings{Name: "test_name"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Test: %s", r.Name)
}
