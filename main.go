package main

import (
	"os"
	"student-api/config"
	"student-api/routers"
	"golang.org/x/exp/slog"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
    config.ConnectDatabase()

    router := routers.SetupRouter()
    router.Run(":8082")
	slog.Info("started")

}