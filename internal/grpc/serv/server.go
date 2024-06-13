package serv

import (
	"context"
	"fmt"
	"github.com/GrosbergKirr/Server_hostname/tools"
	servV1 "github.com/GrosbergKirr/proto_contracts/gen/go/service"
	"google.golang.org/grpc"
	"time"
)

type ServerAPI struct {
	servV1.UnimplementedGatewayServiceServer
}

func Register(grpc *grpc.Server) {
	servV1.RegisterGatewayServiceServer(grpc, &ServerAPI{})
}

func (s *ServerAPI) ChangeHostName(ctx context.Context, req *servV1.HostRequest) (*servV1.HostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ok := make(chan string)

	go func() {
		err := tools.HostNameChanger(req.GetNewHostName(), req.GetPassword(), ok)
		if err != nil {
			fmt.Println("error: ", err)
		}
	}()
	var res string
	// Ожидание завершения операции или таймаута
	select {
	case res = <-ok:
		return &servV1.HostResponse{Result: res}, nil
	case <-ctx.Done():
		return &servV1.HostResponse{Result: "Access denied. Check password"}, nil

	}

}

func (s *ServerAPI) DNSChange(ctx context.Context, req *servV1.DNSRequest) (*servV1.DNSResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ok := make(chan string)

	go func() {
		err := tools.SetDNSServers(req.GetNewDNSName(), req.GetPassword(), ok)
		if err != nil {
			fmt.Println("error: ", err)
		}
	}()
	var res string
	// Ожидание завершения операции или таймаута
	select {
	case res = <-ok:
		return &servV1.DNSResponse{Result: res}, nil
	case <-ctx.Done():
		return &servV1.DNSResponse{Result: "Access denied. Check password"}, nil

	}

}
