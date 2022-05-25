package server

import (
	"context"
	"test_grpc/generated/protos/test"
)

type TestServer struct {
	test.UnimplementedTestServiceServer
}

func NewTestServer() *TestServer {
	return &TestServer{}
}

func (s *TestServer) Hello(ctx context.Context, settings *test.Settings) (*test.Greetings, error) {
	return &test.Greetings{Name: settings.Name}, nil
}
