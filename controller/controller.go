package controller

import "github.com/YuheiTakagawa/tsukupro-server/db"
import (
	"fmt"
	pb "github.com/YuheiTakagawa/tsukupro-server/proto"
	gorp "gopkg.in/gorp.v1"
)

func NewUserController(user *pb.User) (*pb.Status, error) {
	// set data to DB
	dbmap := db.InitDb()
	defer dbmap.Db.Close()

	// checking I have the tx
	fmt.Printf("id: %s, name: %s, birth: %s\n", user.UserId, user.Name, user.Birth)
	if existId(user.UserId, dbmap) {
		return &pb.Status{
			Message: "Duplicate id",
		}, nil
	}

	setdata(user, dbmap)

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
		UserId: 1,
		Type:   2,
		Data:   []byte("ok"),
	}
	list = append(list, proreq)
	proreq2 := &pb.Proreq{
		TxId:   "001",
		UserId: 1,
		Type:   2,
		Data:   []byte("name"),
	}
	list = append(list, proreq2)
	return &pb.ProreqList{
		Req: list,
	}, nil
}

func existId(id int32, dbmap *gorp.DbMap) bool {
	//check id in DB
	var numbers []int32
	_, err := dbmap.Select(&numbers, "select userid from userinfo")
	db.CheckErr(err, "Select failed")
	for _, p := range numbers {
		if id == p {
			return true
		}
	}
	return false
}

func setdata(user *pb.User, dbmap *gorp.DbMap) error {

	u := &db.UserInfo{user.UserId, user.Name, user.Birth}
	err := dbmap.Insert(u)
	db.CheckErr(err, "Insert failed")

	var getuser []db.UserInfo
	_, err = dbmap.Select(&getuser, "select * from userinfo order by userid")
	db.CheckErr(err, "Select failed")
	fmt.Println("All rows:")
	for x, p := range getuser {
		fmt.Printf("	%d: %v\n", x, p)
	}
	return nil
}

func ReflectJudge(judge *pb.Judge) (*pb.Status, error) {
	fmt.Printf("Getting Judge %s\n", judge)
	return &pb.Status{
		Message: "OK",
	}, nil
}
