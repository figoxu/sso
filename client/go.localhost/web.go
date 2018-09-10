package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"net/http"
	"log"
	"gitlab.com/go-box/pongo2gin"
	"github.com/flosch/pongo2"
	"github.com/figoxu/sso/common"
	"github.com/gin-contrib/static"
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
	ssoMid := common.BuildMid(sysEnv.domain, sysEnv.sso_redirect_url, sysEnv.grpc_client, sysEnv.token_cache)
	r.Use(ssoMid, static.Serve("/dist", static.LocalFile("./dist", true)))
	r.HTMLRender = pongo2gin.Default()
	r.Use(ssoMid)
	r.GET("/main", h_main)
	return r
}

func h_main(c *gin.Context) {
	c.HTML(200, "main.html", pongo2.Context{})
}
