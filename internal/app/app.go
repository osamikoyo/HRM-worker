package app

import (
	"context"
	"fmt"
	"net"

	"github.com/osamikoyo/hrm-worker/internal/config"
	"github.com/osamikoyo/hrm-worker/internal/server"
	"github.com/osamikoyo/hrm-worker/pkg/loger"
	"github.com/osamikoyo/hrm-worker/pkg/pb"
	"google.golang.org/grpc"
)

type App struct{
	gRPC *grpc.Server
	loger loger.Logger
	config *config.Config
}

func Init() (*App, error) {
	cfg, err := config.LoadConfig()
	if err != nil{
		return nil, err
	}
	server := grpc.NewServer()
	return &App{
		config: &cfg,
		loger: loger.New(),
		gRPC: server,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	go func ()  {
		<- ctx.Done()
		a.gRPC.Stop()
		a.loger.Info().Msg("Server shutdown!:3")
	}()

	a.loger.Info().Msg("starting the grpc server...")
	pb.RegisterWorkerServiceServer(a.gRPC, &server.Server{})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.config.Host, a.config.Port))
	if err != nil{
		return err
	}

	a.loger.Info().Str("addr", lis.Addr().String()).Msg("Server started!:3")
	return a.gRPC.Serve(lis)
}