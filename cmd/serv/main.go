package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/GrosbergKirr/Server_hostname/internal/app/grpc_app"
	"github.com/GrosbergKirr/Server_hostname/internal/config"
	"github.com/GrosbergKirr/Server_hostname/internal/logger"
	"github.com/GrosbergKirr/Server_hostname/internal/rest"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.SetLogger()

	log.Info("Starting app", slog.Any("cfg", cfg))
	server := grpc_app.NewApp(log, cfg.GRPC.Host, cfg.GRPC.Port)
	sig := make(chan os.Signal)
	go func() {
		err := server.ServerRun()
		if err != nil {
			log.Info("Cant start server")
		}
	}()
	err := rest.RunRest(log)
	if err != nil {
		log.Info("Cant start server", slog.Any("err", err))
	}

	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-sig
	server.ServerStop()

}
