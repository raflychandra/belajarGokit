package belajarGoKit

import (
	"belajarGoKit/proto"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/test/bufconn"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

type server struct{}

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	proto.RegisterAddServiceServer(s, &server{})
	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func (s *server) SearchMovie(ctx context.Context, request *proto.Request) (*proto.Response, error) {
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


func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestSearchMovie(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := proto.NewAddServiceClient(conn)
	resp, err := client.SearchMovie(ctx, &proto.Request{Name: "batman", Page: "1"})
	if err != nil {
		t.Fatalf("error failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
}
