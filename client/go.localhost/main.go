package main

import (
	"github.com/astaxie/beego/config"
	"github.com/quexer/utee"
	"fmt"
	"google.golang.org/grpc"
)

var sysEnv = SysEnv{}

type SysEnv struct {
	domain     string
	grpcClient *grpc.ClientConn
}

func main() {
	fmt.Println("准备启动")
	cfg_core, err := config.NewConfig("ini", "conf.ini")
	utee.Chk(err)
	sysEnv.domain = cfg_core.String("http::domain")
	fmt.Println(sysEnv.domain)
	sysEnv.grpcClient, err = grpc.Dial(cfg_core.String("grpc::server"), grpc.WithInsecure())
	utee.Chk(err)
	defer sysEnv.grpcClient.Close()
	initWeb(cfg_core.String("http::port"))
}

