package main

import (
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc"
	"net"
	"github.com/figoxu/sso/pb/sso"
	"context"
)

var (
	ssoService = &SsoService{}
)

func initGrpcServer(address string) {
	listen, err := net.Listen("tcp", address)
	chk(err)
	s := grpc.NewServer()
	sso.RegisterSsoServiceServer(s, ssoService)
	grpclog.Info("Listen on" + address)
	s.Serve(listen)
}

func chk(err error) {
	if err != nil {
		grpclog.Fatalln(err)
	}
}

type SsoService struct{}

func (p *SsoService) GetLoginInfo(context.Context, *sso.LoginInfoReq) (*sso.LoginInfoRsp, error) {
	return nil, nil
}
func (p *SsoService) SaveUserInfo(context.Context, *sso.User) (*sso.User, error) {
	return nil, nil
}
