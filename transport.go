package belajarGoKit

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

type Search struct {
	Search       []Data `json:"Search,omitempty"`
	TotalResults string `json:"totalResults,omitempty"`
	Response     string `json:"Response"`
	Error        string `json:"Error,omitempty"`
}

type Data struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	IMDBID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func decodeSearchRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	reqName, ok := r.URL.Query()["search"]
	if !ok || len(reqName[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return nil, errors.New("Url Param 'key' is missing")
	}
	reqPage, ok := r.URL.Query()["page"]
	if !ok || len(reqPage[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return nil, errors.New("Url Param 'key' is missing")
	}
	var req searchRequest
	req.Name = reqName[0]
	req.Page = reqPage[0]
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
