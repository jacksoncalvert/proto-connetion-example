package main

import (
	"context"
	"log"
	"net"
	"os"
	"fmt"

	"hello-plugin/proto"
	"google.golang.org/grpc"

	"github.com/hashicorp/go-hclog"
)


type Server struct {
	proto.UnimplementedHelloPluginServer
}

func main() {
	CreateServer()
	fmt.Print("helloExample")
}

func (s *Server) GetPerson(ctx context.Context, req *proto.RequestPerson) (*proto.Person, error) {
        return &proto.Person{Name: "Jackson Lane"}, nil
}


func CreateServer() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Debug,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	lis, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloPluginServer(s, &Server{})
	logger.Debug("Plugin initialised")
	s.Serve(lis)

}
