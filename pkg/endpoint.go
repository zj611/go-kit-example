package pkg

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Reply string `json:"reply"`
}


type ByeRequest struct {
	//Name string `json:"name"`
	A int `json:"a"`
	B int `json:"b"`
}

type ByeResponse struct {
	Reply string `json:"reply"`
}

func MakeServerEndpointHello(s MyServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r,ok := request.(HelloRequest)
		if !ok {
			return HelloResponse{},nil
		}
		return HelloResponse{Reply: s.Hello(r.Name)},nil
	}
}

func MakeServerEndpointBye(s MyServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r,ok := request.(ByeRequest)
		if !ok {
			return ByeResponse{},nil
		}
		return ByeResponse{Reply: s.Bye(r.A, r.B)},nil
	}
}

