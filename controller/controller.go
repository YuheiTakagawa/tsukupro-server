package controller

import (
	pb "../proto"
	"fmt"
)

func NewUserController(user *pb.User) (*pb.Status, error) {
	// checking I have the tx
	fmt.Printf("id: %s, name: %s, birth: %s\n", user.UserId, user.Name, user.Birth)
	if err := checkId(user.UserId); err != nil {
		return &pb.Status{
			Message: "Duplicate id",
		}, nil
	}

	setdata(user)

	return &pb.Status{
		Message: "OK",
	}, nil
}

func EditProfController(proreq *pb.Proreq) (*pb.Status, error) {
	return &pb.Status{
		Message: "OK",
	}, nil
}

type Proreqs []*pb.Proreq

func SearchProfController(id *pb.UserId) (*pb.ProreqList, error) {
	var list Proreqs
	proreq := &pb.Proreq{
		TxId:   "000",
		UserId: "000",
		Type:   2,
		Data:   []byte("ok"),
	}
	list = append(list, proreq)
	proreq.UserId = "roto"
	list = append(list, proreq)
	return &pb.ProreqList{
		Req: list,
	}, nil
}

func checkId(id string) error {
	//check id in DB
	return nil
}

func setdata(user *pb.User) error {
	// set data to DB
	return nil
}
