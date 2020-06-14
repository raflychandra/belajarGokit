package main

import (
	"belajarGoKit/proto"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &Server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *Server) SearchMovie(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	var search proto.Response
	name, page := request.GetName(), request.GetPage()
	str := fmt.Sprintf("http://omdbapi.com/?apikey=faf7e5bb&s=%s&page=%s", name, page)
	response, err := http.Get(str)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(data, &search)

	return &proto.Response{
		Search:       search.Search,
		TotalResults: search.TotalResults,
		Response:     search.Response,
		Error:        search.Error,
	}, nil
}
