package models

import (
	"github.com/osamikoyo/hrm-worker/pkg/pb"
)

type Worker struct{
	UserID uint64  `gorm:"primaryKey;autoIncrement"`
	Firstname string 
	Secondname string
	Salary uint64
	Email string
	Post string
}

func ToModels(pbwork *pb.Worker) *Worker {
	work := &Worker{
		UserID: pbwork.UserID,
		Salary: pbwork.Salary,
		Firstname: pbwork.Firstname,
		Secondname: pbwork.Secondname,
		Post: pbwork.Post,
		Email: pbwork.Email,
	}

	return work
}

func ToPB(worker *Worker) *pb.Worker {
	work := &pb.Worker{
		UserID: worker.UserID,
		Email: worker.Email,
		Salary: worker.Salary,
		Firstname: worker.Firstname,
		Secondname: worker.Secondname,
		Post: worker.Post,
	}
	return work
}