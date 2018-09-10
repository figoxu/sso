package main

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/figoxu/gh"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"fmt"
	"github.com/figoxu/sso/common"
	"net/url"
	jsonp "github.com/jim3ma/gin-jsonp"
)

func initWeb(port string) {
	engine := mount()
	http.Handle("/", engine)
	log.Fatal(http.ListenAndServe(port, nil))
}

func mount() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	store := cookie.NewStore([]byte("xujh945@qq.com"))
	r.Use(sessions.Sessions("figoxu", store))
	sso := r.Group("/sso", m_gh)
	{
		login := sso.Group("/login")
		{
			login.GET("/redirect", h_redirect)
			login.POST("/exc", h_login)
			login.GET("/token", jsonp.Handler(), h_check_login)
		}
	}
	return r
}

//curl http://sso.localhost/sso/login/redirect?from=https%3a%2f%2fgithub.com%2ffigoxu%2fsso
func h_redirect(ctx *gin.Context) {
	saveTokenFromAddress(ctx)
	jumpLoc := jumpUrl(ctx)
	ctx.Redirect(http.StatusFound, jumpLoc)
	return
}

//curl http://sso.localhost/sso/login/exc -d "username=figo&password=hello"
func h_login(c *gin.Context) {
	type LoginResp struct {
		SuccessFlag bool
		JumpUrl     string
	}
	resp := LoginResp{}
	env := c.MustGet("env").(*Env)
	fh := env.fh
	username, password := fh.String("username"), fh.String("password")
	validate := func() (bool, string) {
		if username == "" || password == "" {
			return false, ""
		}
		user := NewUserDao(pg_rbac).GetByLoginName(username)
		passwordSaltHelp := NewUserPasswordSaltHelper(user)
		if user.Id <= 0 || password != passwordSaltHelp.Decode(user.Password) {
			return false, ""
		}
		th := NewTokenHelper(c)
		basicPureToken := th.NewToken(user.Id)
		fmt.Println(">>>>>>   Write @Domain:", sysEnv.cookie_domain)
		common.WriteCookie(c, basicPureToken, sysEnv.cookie_domain)
		common.StoreToken2Session(c, basicPureToken)
		return true, basicPureToken
	}
	successFlag, _ := validate()
	if !successFlag {
		resp.SuccessFlag = false
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.SuccessFlag = true
	redirectUrl := jumpUrl(c)
	resp.JumpUrl = url.QueryEscape(redirectUrl)
	c.JSON(http.StatusOK, resp)
}

//curl http://sso.localhost/sso/login/token?from=https%3a%2f%2fgithub.com%2ffigoxu%2fsso
func h_check_login(ctx *gin.Context) {
	loginFlag := checkLogin(ctx)
	jumpLoc := jumpUrl(ctx)
	basicPureToken := common.GetBasicPureToken(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"loginFlag":      loginFlag,
		"jumpLoc":        jumpLoc,
		"basicPureToken": basicPureToken,
	})
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
