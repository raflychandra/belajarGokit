package belajarGoKit

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SearchEndpoint   endpoint.Endpoint
}

func ControllerSearch(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(searchRequest)
		b, err := srv.Search(ctx, req.Name, req.Page)
		if err != nil {
			return b, err
		}
		return b, nil
	}
}