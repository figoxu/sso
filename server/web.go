package main

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/figoxu/gh"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"fmt"
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
	env := ctx.MustGet("env").(*Env)
	from := env.ph.String("from")
	saveFromAddress := func() {
		session := sessions.Default(ctx)
		session.Set(SSO_FROM, from)
	}
	checkLogin := func() bool {
		basicRawToken, err := ctx.Cookie(SSO_TOKEN_COOKIE)
		if err != nil {
			return false
		}
		if basicRawToken == "" {
			session := sessions.Default(ctx)
			if brt := session.Get(SSO_BASIC_RAW_TOKEN); brt == nil {
				return false
			} else {
				basicRawToken = fmt.Sprint(brt)
			}
		}
		tokenHelper := NewTokenHelper(ctx)
		uid, rawToken := tokenHelper.ParseToken(basicRawToken)
		return tokenHelper.CheckRawToken(uid, rawToken)
	}
	saveFromAddress()
	jumpLoc := from
	if checkLogin() && jumpLoc == "" {
		jumpLoc = sysEnv.welcome_page
	} else {
		jumpLoc = sysEnv.login_page
	}
	ctx.Redirect(http.StatusOK, sysEnv.welcome_page)
	return
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
	validate := func() bool {
		if username == "" || password == "" {
			return false
		}
		user := NewUserDao(pg_rbac).GetByLoginName(username)
		passwordSaltHelp := NewUserPasswordSaltHelper(user)
		if user.Id <= 0 || user.Password != passwordSaltHelp.Decode(password) {
			return false;
		}
		return true
	}
	if !validate() {
		resp.SuccessFlag = false
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.SuccessFlag = true
	session := sessions.Default(c)
	redirect_address := session.Get(SSO_FROM)
	if redirect_address == nil {
		resp.JumpUrl = sysEnv.welcome_page
	} else {
		resp.JumpUrl = fmt.Sprint(redirect_address)
	}
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
