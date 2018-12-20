package handler

import (
	"context"

	"github.com/YuheiTakagawa/tsukupro-server/controller"
	pb "github.com/YuheiTakagawa/tsukupro-server/proto"
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

func (s *Tsukupro) SearchProf(id *pb.UserId, stream pb.Tsukupro_SearchProfServer) error {
	err := controller.SearchProfController(id, stream)
	return err
}

func (s *Tsukupro) SendJudge(ctx context.Context, judge *pb.Judge) (*pb.Status, error) {
	ret, err := controller.ReflectJudge(judge)
	return ret, err
}
