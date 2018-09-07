package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"net/http"
	"log"
)

func initWeb(port string) {
	engine := mount()
	http.Handle("/", engine)
	log.Fatal(http.ListenAndServe(port, nil))
}

func mount() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("xujh945@qq.com"))
	r.Use(sessions.Sessions("figoxu", store))
	return r
}