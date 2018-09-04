package main

import (
	"github.com/astaxie/beego/config"
	"github.com/quexer/utee"
)

func main() {
	cfg_core, err := config.NewConfig("ini", "conf.ini")
	utee.Chk(err)
	initDb(cfg_core.String("db_sso::param"))
	initGrpcServer(cfg_core.String("grpc::address"))
	initWeb(cfg_core.String("http::port"))
}
