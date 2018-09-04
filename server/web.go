package main

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/figoxu/gh"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
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
	sso := r.Group("/sso", m_gh)
	{
		login := sso.Group("/login")
		{
			login.GET("/redirect/:from", h_redirect)
			login.POST("/exc", h_login)
		}
	}
	return r
}

func h_redirect(ctx *gin.Context) {
	//todo 保存 from的地址到 session

	//todo 验证 cookie 是否包含已登陆信息【已登陆，则跳回之前的地址】

	//todo 未登录，则调到登陆页面
}

func h_login(c *gin.Context) {
	type LoginResp struct {
		SuccessFlag bool
		JumpUrl     string
	}
	resp := LoginResp{}
	env := c.MustGet("env").(*Env)
	fh := env.fh
	username, password := fh.String("username"), fh.String("password")
	validate := func() bool{
		if username == "" || password == "" {
			return false
		}
		user:=NewUserDao(pg_rbac).GetByLoginName(username)
		if user.Id<=0 || user.Password!=password {
			//todo Password需要加密

		}
		return true
	}
	//todo 对登陆信息验证
	if  !validate(){
		resp.SuccessFlag = false
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.SuccessFlag = true
	//resp.JumpUrl =
	c.JSON(http.StatusOK, resp)
}

type Env struct {
	fh *gh.FormHelper
	ph *gh.ParamHelper
}

func m_gh(c *gin.Context) {
	c.Set("env", &Env{
		fh: gh.NewFormHelper(c),
		ph: gh.NewParamHelper(c),
	})
	c.Next()
}
