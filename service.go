package belajarGoKit

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service interface {
	Search(ctx context.Context, name, page string) (interface{}, error)
}

type dateService struct{}

// NewService makes a new Service.
func NewService() Service {
	return dateService{}
}

func (dateService) Search(ctx context.Context, name, page string) (interface{}, error) {
	var search Search
	str := fmt.Sprintf("http://omdbapi.com/?apikey=faf7e5bb&s=%s&page=%s", name, page)
	response, err := http.Get(str)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return "", err
	}
	data, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(data, &search)
	return search, nil
}