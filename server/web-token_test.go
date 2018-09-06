package main

import (
	"testing"
	"log"
	"github.com/figoxu/Figo"
)

func TestNewSaltHelper(t *testing.T) {
	sh := NewSaltHelper("hello", "world")
	content := "figo'xu awesome"
	result := sh.Encode(content)
	log.Println(result)
	val := sh.Decode(result)
	log.Println(val)
	log.Println(val == content)
}

func TestBasicAuthDecode(t *testing.T) {
	s := "hello#world,how_are_you"
	content := BasicAuthEncode(s)
	log.Println(content)
	result := BasicAuthDecode(content)
	log.Println(result)
	log.Println(result == s)
}

func TestNewUserPasswordSaltHelper(t *testing.T) {
	initTestEnv()
	userDao:=NewUserDao(pg_rbac)
	u:=userDao.GetByLoginName("figo")
	log.Println(Figo.JsonString(u))
	saltHelper:=NewUserPasswordSaltHelper(u)
	pwd:=saltHelper.Decode(u.Password)
	log.Println("PASSWORD IS : ",pwd)
}
