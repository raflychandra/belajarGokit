package server

import (
	"belajarGoKit/proto"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SearchEndpoint   endpoint.Endpoint
}

func ControllerSearch(srv proto.AddServiceClient) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(searchRequest)
		payload := &proto.Request{
			Name: req.Name,
			Page: req.Page,
		}
		b, err := srv.SearchMovie(ctx, payload)
		if err != nil {
			return b, err
		}
		return b, nil
	}
}