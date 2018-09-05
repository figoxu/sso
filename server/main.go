package main

import (
	"github.com/astaxie/beego/config"
	"github.com/quexer/utee"
)

var sysEnv = SysEnv{}

type SysEnv struct {
	domain       string
	welcome_page string
	login_page   string
}

func main() {
	cfg_core, err := config.NewConfig("ini", "conf.ini")
	utee.Chk(err)
	sysEnv.domain = cfg_core.String("http::domain")
	sysEnv.welcome_page = cfg_core.String("http::welcome_page")
	sysEnv.login_page = cfg_core.String("http::login_page")
	initDb(cfg_core.String("db_sso::param"))
	initGrpcServer(cfg_core.String("grpc::address"))
	initWeb(cfg_core.String("http::port"))
}
