package transport

import (
	"fmt"

	"github.com/osamikoyo/hrm-worker/internal/config"
	"github.com/osamikoyo/hrm-worker/pkg/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct{
	Client pb.WorkerServiceClient
}

func InitServiceClient(c *config.Config) pb.WorkerServiceClient {
	cc, err := grpc.Dial(c.Host, grpc.WithInsecure())

    if err != nil {
        fmt.Println("Could not connect:", err)
    }

	return pb.NewWorkerServiceClient(cc)
}