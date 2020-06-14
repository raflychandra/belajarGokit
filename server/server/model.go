package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type searchRequest struct {
	Name string
	Page string
}

func decodeSearchRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	reqName, ok := r.URL.Query()["search"]
	if !ok || len(reqName[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return nil, errors.New("Url Param 'name' is missing")
	}
	reqPage, ok := r.URL.Query()["page"]
	if !ok || len(reqPage[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return nil, errors.New("Url Param 'page' is missing")
	}
	var req searchRequest
	req.Name = reqName[0]
	req.Page = reqPage[0]
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
