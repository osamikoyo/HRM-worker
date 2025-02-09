package server

import (
	"context"
	"net/http"

	"github.com/osamikoyo/hrm-worker/internal/data"
	"github.com/osamikoyo/hrm-worker/internal/data/models"
	"github.com/osamikoyo/hrm-worker/pkg/pb"
)

type Server struct{
	Storage *data.Storage
	pb.UnimplementedWorkerServiceServer
}

func (s *Server) Add(_ context.Context, req *pb.AddWorkerRequest) (*pb.AddWorkerResponse, error){
	id, err := s.Storage.Create(models.ToModels(req.Worker))
	return &pb.AddWorkerResponse{
		Respone: &pb.Response{
			Error: err.Error(),
			Status: http.StatusOK,
		},
		UserID: id,
	}, err
}

func (s *Server) Delete(_ context.Context, req *pb.DeleteWorkerRequest) (*pb.Response, error){
	err := s.Storage.Delete(req.UserID)
	return &pb.Response{
		Error: err.Error(),
		Status: http.StatusOK,
	}, err
}

func (s *Server) Get(_ context.Context, req *pb.GetWorkerRequest) (*pb.GetWorkerResponse, error){
	resp, err := s.Storage.Get(req.UserID)
	return &pb.GetWorkerResponse{
		Response: &pb.Response{
			Error: err.Error(),
			Status: http.StatusOK,
		},
		Worker: models.ToPB(resp),
	}, err
}

func (s *Server) Update(_ context.Context,req *pb.UpdateWorkerRequest) (*pb.Response, error){
	err := s.Storage.Update(req.UserID, models.ToModels(req.NewWorkerParametres))
	return &pb.Response{
		Error: err.Error(),
		Status: http.StatusOK,
	}, err
}