package main

import (
	"testing"
	"github.com/quexer/utee"
	"github.com/astaxie/beego/config"
	"log"
	"github.com/figoxu/Figo"
)

func initTestEnv(){
	cfg_core, err := config.NewConfig("ini", "conf.ini")
	utee.Chk(err)
	initDb(cfg_core.String("db_sso::param"))
}

func TestUserDao_Save(t *testing.T) {
	initTestEnv()
	u:=User{
		Account:"figo",
		Password:"123456",
		PasswordSalt:"figoxu",
	}
	sh:=NewUserPasswordSaltHelper(u)
	u.Password = sh.Encode(u.Password)
	userDao:= NewUserDao(pg_rbac)
	userDao.Save(&u)
	log.Println(Figo.JsonString(u))
}

