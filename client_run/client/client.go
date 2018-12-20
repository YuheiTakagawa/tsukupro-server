package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/YuheiTakagawa/tsukupro-server/proto"
)

type ClientInfo struct {
	Conn proto.TsukuproClient
	Id   string
	User *proto.User
}

func (c *ClientInfo) Search() error {
	message := &proto.UserId{
		Id: c.User.UserId,
	}

	var list []*proto.Proreq
	stream, _ := c.Conn.SearchProf(context.TODO(), message)
	for {
		proreq, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Get all profs")
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
		list = append(list, proreq)
	}
	for _, r := range list {
		fmt.Printf("result: %s\n", r)
		c.Judge(r.TxId)
		fmt.Println("")
	}
	return nil
}

func (c *ClientInfo) Judgement(txid string, judgement bool) *proto.Judge {
	message := &proto.Judge{
		UserId: c.User.UserId,
		TxId:   txid,
		Res:    judgement,
	}
	fmt.Printf("judge: %s\n", message)

	return message
}

func (c *ClientInfo) Judge(txid string) error {
	judgement := false
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("Judge this tx id=%s  \"true\" or \"false\"?\n", txid)
	stdin.Scan()
	if jstr := stdin.Text(); jstr == "true" {
		judgement = true
	}
	judge := c.Judgement(txid, judgement)
	res, _ := c.Conn.SendJudge(context.TODO(), judge)
	fmt.Printf("result: %s\n", res)
	return nil
}

func (c *ClientInfo) Create() error {
	fmt.Printf("user info %s\n", c.User)
	res, _ := c.Conn.NewUser(context.TODO(), c.User)
	fmt.Printf("result: %s\n", res)
	return nil
}
