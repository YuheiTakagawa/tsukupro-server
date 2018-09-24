package handler

import (
	"../controller"
	pb "../proto"
	"context"
)

type Tsukupro struct{}

func (s *Tsukupro) NewUser(ctx context.Context, user *pb.User) (*pb.Status, error) {
	if ret, err := controller.NewUserController(user); err != nil {
		return &pb.Status{
			Message: "DENY",
		}, nil
	} else {
		return ret, nil
	}
}

func (s *Tsukupro) EditProf(ctx context.Context, proreq *pb.Proreq) (*pb.Status, error) {
	if ret, err := controller.EditProfController(proreq); err != nil {
		return &pb.Status{
			Message: "DENY",
		}, nil
	} else {
		return ret, nil
	}
}

func (s *Tsukupro) SearchProf(ctx context.Context, id *pb.UserId) (*pb.ProreqList, error) {
	ret, _ := controller.SearchProfController(id)
	return ret, nil
}
