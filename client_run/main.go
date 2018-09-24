package main

import "github.com/YuheiTakagawa/tsukupro-server/proto"
import "github.com/YuheiTakagawa/tsukupro-server/client_run/client"
import (
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("usage: ./client <IP> <command>")
		os.Exit(1)
	}
	fmt.Printf("%s\n", os.Args[1])
	command := os.Args[2]

	conn, err := grpc.Dial(os.Args[1]+":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	cli := proto.NewTsukuproClient(conn)

	userinfo := &proto.User{
		UserId: "000",
		Name:   "Yuhei",
		Birth:  "1995-11-27",
	}
	c := &client.ClientInfo{
		Conn: cli,
		Id:   "im",
		User: userinfo,
	}

	switch command {
	case "search":
		c.Search()
	case "new":
		c.Create()
	}
}
