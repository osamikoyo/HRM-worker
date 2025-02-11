package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/osamikoyo/hrm-worker/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app, err := app.Init()
	if err != nil{
		fmt.Println(err)
	}

	if err = app.Run(ctx);err != nil{
		fmt.Println(err)
	}
}