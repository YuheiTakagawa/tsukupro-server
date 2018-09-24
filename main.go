package main

import "github.com/YuheiTakagawa/tsukupro-server/handler"
import "github.com/YuheiTakagawa/tsukupro-server/proto"
import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const grpcPort = ":50051"

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	server := &handler.Tsukupro{}
	proto.RegisterTsukuproServer(s, server)
	fmt.Println("server start")
	s.Serve(lis)
}
