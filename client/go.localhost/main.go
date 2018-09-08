package main

import (
	"github.com/astaxie/beego/config"
	"github.com/quexer/utee"
	"fmt"
	"google.golang.org/grpc"
)

var sysEnv = SysEnv{}

type SysEnv struct {
	domain           string
	sso_redirect_url string
	grpc_client      *grpc.ClientConn
	token_cache      *utee.TimerCache
}

func main() {
	fmt.Println("准备启动")
	cfg_core, err := config.NewConfig("ini", "conf.ini")
	utee.Chk(err)
	sysEnv.token_cache = utee.NewTimerCache(60, nil)
	sysEnv.domain = cfg_core.String("http::domain")
	sysEnv.sso_redirect_url = cfg_core.String("sso::redirect_url")
	fmt.Println(sysEnv.domain)
	sysEnv.grpc_client, err = grpc.Dial(cfg_core.String("grpc::server"), grpc.WithInsecure())
	utee.Chk(err)
	defer sysEnv.grpc_client.Close()
	initWeb(cfg_core.String("http::port"))
}
