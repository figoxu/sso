package main

import (
	"github.com/astaxie/beego/config"
	"github.com/quexer/utee"
	"fmt"
)

var sysEnv = SysEnv{}

type SysEnv struct {
	cookie_domain       string
	welcome_page string
	login_page   string
}

func main() {
	fmt.Println("准备启动")
	cfg_core, err := config.NewConfig("ini", "conf.ini")
	utee.Chk(err)
	sysEnv.cookie_domain = cfg_core.String("http::cookie_domain")
	fmt.Println(sysEnv.cookie_domain)
	sysEnv.welcome_page = cfg_core.String("http::welcome_page")
	sysEnv.login_page = cfg_core.String("http::login_page")
	fmt.Println("->数据库")
	initDb(cfg_core.String("db_sso::param"))
	fmt.Println("->GRPC")
	go initGrpcServer(cfg_core.String("grpc::address"))
	fmt.Println("->Web")
	initWeb(cfg_core.String("http::port"))
}
