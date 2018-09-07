package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"net/http"
	"log"
	"gitlab.com/go-box/pongo2gin"
	"github.com/flosch/pongo2"
)

func initWeb(port string) {
	engine := mount()
	http.Handle("/", engine)
	log.Fatal(http.ListenAndServe(port, nil))
}

func mount() *gin.Engine {
	r := gin.Default()
	r.HTMLRender = pongo2gin.Default()
	store := cookie.NewStore([]byte("xujh945@qq.com"))
	r.Use(sessions.Sessions("figoxu", store))
	r.GET("/main",h_main)
	return r
}

func h_main(c *gin.Context){
	c.HTML(200, "main.html", pongo2.Context{})
}



