package grpc_app

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/GrosbergKirr/Server_hostname/internal/grpc/serv"
	"google.golang.org/grpc"
)

type ServerStrucr struct {
	log    *slog.Logger
	Server *grpc.Server
	Host   string
	Port   int
}

func NewApp(log *slog.Logger, host string, port int) *ServerStrucr {
	Server := grpc.NewServer()
	serv.Register(Server)
	return &ServerStrucr{log, Server, host, port}
}

func (a *ServerStrucr) ServerRun() error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.Host, a.Port))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	a.log.Info("Starting Server on", slog.String("address", l.Addr().String()))
	err = a.Server.Serve(l)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (a *ServerStrucr) ServerStop() {
	a.log.Info("Server is stopped")
	a.Server.GracefulStop()
}
