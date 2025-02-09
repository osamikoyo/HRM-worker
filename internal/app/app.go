package app

import (
	"github.com/osamikoyo/hrm-worker/internal/config"
	"github.com/osamikoyo/hrm-worker/pkg/loger"
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
}