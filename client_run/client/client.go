package client

import "../../proto"
import (
	"context"
	"fmt"
)

type ClientInfo struct {
	Conn proto.TsukuproClient
	Id string
	User *proto.User
}

func (c *ClientInfo) Search() (error) {
	message := &proto.UserId{
		Id: c.Id,
	}

	res, _ := c.Conn.SearchProf(context.TODO(), message)
	fmt.Printf("result: %s\n", res)
	c.Judge("000")
	return nil
}

func (c *ClientInfo) Judgement(txid string, res bool) (error) {
	message := &proto.Judge{
		UserId: c.Id,
		TxId: "000",
		Res: res,
	}
	fmt.Printf("judge: %s\n", message)

	return nil
}

func (c *ClientInfo) Judge(txid string) (error) {
	c.Judgement(txid, true)
	return nil
}

func (c *ClientInfo) Create() (error) {
	fmt.Printf("user info %s\n", c.User)
	res, _ := c.Conn.NewUser(context.TODO(), c.User)
	fmt.Printf("result: %s\n", res)
	return nil
}
