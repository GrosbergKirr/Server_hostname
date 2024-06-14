package main

import (
	"github.com/GrosbergKirr/Server_hostname/internal/app/grpc_app"
	"github.com/GrosbergKirr/Server_hostname/internal/config"
	"github.com/GrosbergKirr/Server_hostname/internal/logger"
	"github.com/GrosbergKirr/Server_hostname/internal/rest"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Set config and logger
	cfg := config.LoadConfig()
	log := logger.SetLogger()

	//Start gRPC-Server
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
	//Graceful shutdown
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-sig
	server.ServerStop()

}
