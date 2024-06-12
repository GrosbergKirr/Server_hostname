package serv

import (
	"context"
	"fmt"
	"github.com/GrosbergKirr/Server_hostname/tools"
	servV1 "github.com/GrosbergKirr/proto_contracts/gen/go/service"
	"google.golang.org/grpc"
)

type ServerAPI struct {
	servV1.UnimplementedGatewayServiceServer
}

func Register(grpc *grpc.Server) {
	servV1.RegisterGatewayServiceServer(grpc, &ServerAPI{})
}

func (s *ServerAPI) ChangeHostName(ctx context.Context, req *servV1.HostRequest) (*servV1.HostResponse, error) {
	res := tools.HostNameChanger(req.GetNewHostName(), req.GetPassword())
	return &servV1.HostResponse{Result: fmt.Sprintf("Name setting success. New name is %s", res)}, nil

}

func (s *ServerAPI) DNSChange(ctx context.Context, req *servV1.DNSRequest) (*servV1.DNSResponse, error) {

	panic("im not written")

}
