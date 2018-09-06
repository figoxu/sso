package main

import (
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc"
	"net"
	"github.com/figoxu/sso/pb/sso"
	"context"
	"github.com/pkg/errors"
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

func (p *SsoService) GetLoginInfo(ctx context.Context, req *sso.LoginInfoReq) (rsp *sso.LoginInfoRsp, err error) {
	if req.BasicRawToken == "" {
		return nil, errors.New("bad param")
	}

	uid, rawToken := ParseToken(req.BasicRawToken)
	if CheckPureToken(uid, rawToken) {
		return nil, errors.New("not auth")
	}
	userDao, resourceDao, userGroupDao := NewUserDao(pg_rbac), NewResourceDao(pg_rbac), NewUserGroupDao(pg_rbac)
	user := userDao.GetById(uid)
	resources := resourceDao.FindByUid(uid)
	userGroups := userGroupDao.FindByUid(uid)

	rs := make([]*sso.Resource, 0)
	for _, resource := range resources {
		rs = append(rs, resource.toProto())
	}
	ugs := make([]*sso.UserGroup, 0)
	for _, ug := range userGroups {
		ugs = append(ugs, ug.toProto())
	}
	rsp = &sso.LoginInfoRsp{
		User:        user.toProto(),
		Resources:   rs,
		UserGroupes: ugs,
	}
	return rsp, nil
}
func (p *SsoService) SaveUserInfo(ctx context.Context, user *sso.User) (userRsp *sso.User, err error) {
	return nil, nil
}
