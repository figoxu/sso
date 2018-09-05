package main

import (
	"github.com/astaxie/beego/config"
	"github.com/quexer/utee"
)

var sysEnv = SysEnv{}

type SysEnv struct {
	domain string
}

func main() {
	cfg_core, err := config.NewConfig("ini", "conf.ini")
	utee.Chk(err)
	sysEnv.domain = cfg_core.String("http::domain")
	initDb(cfg_core.String("db_sso::param"))
	initGrpcServer(cfg_core.String("grpc::address"))
	initWeb(cfg_core.String("http::port"))
}
