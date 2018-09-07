package main

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/figoxu/gh"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"fmt"
	"github.com/figoxu/Figo"
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
			login.GET("/redirect", h_redirect)
			login.POST("/exc", h_login)
		}
	}
	return r
}

//curl http://sso.localhost/sso/login/redirect?from=https%3a%2f%2fgithub.com%2ffigoxu%2fsso
func h_redirect(ctx *gin.Context) {
	vs := ctx.Request.URL.Query()
	from := vs.Get("from")
	fmt.Println(">>>>>>>>>")
	fmt.Println(from)
	fmt.Println("<<<<<<<<<")
	saveFromAddress := func() {
		session := sessions.Default(ctx)
		session.Set(SSO_FROM, from)
		session.Save();
	}
	checkLogin := func() bool {
		basicPureToken, err := ctx.Cookie(SSO_TOKEN_COOKIE)
		if err != nil {
			return false
		}
		if basicPureToken == "" {
			session := sessions.Default(ctx)
			if brt := session.Get(SSO_BASIC_PURE_TOKEN); brt == nil {
				return false
			} else {
				basicPureToken = fmt.Sprint(brt)
			}
		}
		uid, pureToken := ParseToken(basicPureToken)
		return CheckPureToken(uid, pureToken)
	}
	saveFromAddress()
	jumpLoc := sysEnv.login_page
	if checkLogin() {
		if from != "" {
			jumpLoc = from
		} else {
			jumpLoc = sysEnv.welcome_page
		}
	}
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
	validate := func() (bool,string) {
		if username == "" || password == "" {
			return false,""
		}
		user := NewUserDao(pg_rbac).GetByLoginName(username)
		passwordSaltHelp := NewUserPasswordSaltHelper(user)
		if user.Id <= 0 || password != passwordSaltHelp.Decode(user.Password) {
			return false,""
		}
		th := NewTokenHelper(c)
		basicPureToken:=th.NewToken(user.Id)
		return true,basicPureToken
	}
	successFlag,basicPureToken:=validate()
	if !successFlag {
		resp.SuccessFlag = false
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.SuccessFlag = true
	resp.JumpUrl = jumpUrl(sessions.Default(c),basicPureToken)
	c.JSON(http.StatusOK, resp)
}

func jumpUrl(session sessions.Session,basicPureToken string) (redirect_address string) {
	v := session.Get(SSO_FROM)
	if v == nil {
		redirect_address = sysEnv.welcome_page
	} else {
		redirect_address = fmt.Sprint(v)
		redirect_address = Figo.UrlAppendParam(redirect_address,"basic_pure_token",basicPureToken)
	}
	return redirect_address
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
