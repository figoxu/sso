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
	"github.com/figoxu/sso/common"
	"github.com/quexer/utee"
	"net/url"
)

func initWeb(port string) {
	engine := mount()
	http.Handle("/", engine)
	log.Fatal(http.ListenAndServe(port, nil))
}

func mount() *gin.Engine {
	r := gin.Default()
	r.Use(func(*gin.Context){
		defer Figo.Catch()
	})
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
	from := vs.Get(common.SSO_FROM_PARAM)
	fmt.Println(">>>>>>>>>")
	fmt.Println(from)
	fmt.Println("<<<<<<<<<")
	saveFromAddress := func() {
		session := sessions.Default(ctx)
		session.Set(SSO_FROM, from)
		session.Save();
	}
	basicPureToken := common.GetBasicPureToken(ctx)
	checkLogin := func() bool {
		if basicPureToken == "" {
			return false;
		}
		fmt.Println("@BAISC_PURE_TOKEN:", basicPureToken)
		uid, pureToken := ParseToken(basicPureToken)
		return CheckPureToken(uid, pureToken)
	}
	saveFromAddress()
	jumpLoc := sysEnv.login_page
	if checkLogin() {
		fmt.Println("@登陆成功")
		if from != "" {
			from, _ = url.QueryUnescape(from)
			jumpLoc = Figo.UrlAppendParam(from, common.SSO_TOKEN_PARAM, basicPureToken)
		} else {
			jumpLoc = sysEnv.welcome_page
		}
	} else {
		fmt.Println("@自动登陆失败")
	}
	fmt.Println(">>>>>  跳转到服务：", jumpLoc)
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
	successFlag, basicPureToken := validate()
	if !successFlag {
		resp.SuccessFlag = false
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.SuccessFlag = true
	redirectUrl := jumpUrl(sessions.Default(c), basicPureToken)
	resp.JumpUrl = url.QueryEscape(redirectUrl)
	c.JSON(http.StatusOK, resp)
}

func jumpUrl(session sessions.Session, basicPureToken string) (redirect_address string) {
	v := session.Get(SSO_FROM)
	if v == nil {
		return sysEnv.welcome_page
	}
	redirect_address = fmt.Sprint(v)
	redirect_address, err := url.QueryUnescape(redirect_address)
	utee.Chk(err)
	log.Println("@Redirect_Adress:", redirect_address, " @PARAM:", common.SSO_TOKEN_PARAM, " @V:", basicPureToken)
	redirect_address = Figo.UrlAppendParam(redirect_address, common.SSO_TOKEN_PARAM, basicPureToken)
	log.Println("@RESULT_AFTER_APPEND: ", redirect_address)
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
