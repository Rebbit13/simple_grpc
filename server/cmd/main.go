package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"test_grpc/generated/protos/test"
	"test_grpc/server/internal/server"
)

func main() {
	handler := server.NewTestServer()
	grpcServer := grpc.NewServer()
	test.RegisterTestServiceServer(grpcServer, handler)

	lis, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		panic(err)
	}
	log.Printf("Start on address: %s", lis.Addr())
	if err = grpcServer.Serve(lis); err != nil {
		log.Println(err)
	}
}
