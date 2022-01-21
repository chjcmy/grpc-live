package handler

import (
	"context"
	"grpc-example/helloworld"
)

type GreeterHandler struct {
}

func (h GreeterHandler) SayHello(ctx context.Context, req *helloworld.SayHelloRequest) (*helloworld.SayHelloResponse, error) {
	return &helloworld.SayHelloResponse{
		ResponseCode:    200,
		ResponseMessage: "Hi, Nice to meet you, too",
	}, nil
}
