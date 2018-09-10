package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/figoxu/sso/common"
	"github.com/gin-contrib/sessions"
	"github.com/quexer/utee"
	"github.com/figoxu/Figo"
	"net/url"
	"log"
)

func saveTokenFromAddress(ctx *gin.Context){
	vs := ctx.Request.URL.Query()
	from := vs.Get(common.SSO_FROM_PARAM)
	fmt.Println(">>>>>>>>>")
	fmt.Println(from)
	fmt.Println("<<<<<<<<<")
	session := sessions.Default(ctx)
	session.Set(SSO_FROM, from)
	session.Save();
}

func checkLogin(ctx *gin.Context) bool{
	basicPureToken := common.GetBasicPureToken(ctx)
	if basicPureToken == "" {
		return false;
	}
	fmt.Println("@BAISC_PURE_TOKEN:", basicPureToken)
	uid, pureToken := ParseToken(basicPureToken)
	return CheckPureToken(uid, pureToken)
}


func jumpUrl(ctx *gin.Context) (redirect_address string) {
	if(!checkLogin(ctx)){
		return sysEnv.login_page
	}
	session := sessions.Default(ctx)
	v := session.Get(SSO_FROM)
	if v == nil {
		return sysEnv.welcome_page
	}
	redirect_address = fmt.Sprint(v)
	redirect_address, err := url.QueryUnescape(redirect_address)
	utee.Chk(err)
	basicPureToken := common.GetBasicPureToken(ctx)
	log.Println("@Redirect_Adress:", redirect_address, " @PARAM:", common.SSO_TOKEN_PARAM, " @V:", basicPureToken)
	redirect_address = Figo.UrlAppendParam(redirect_address, common.SSO_TOKEN_PARAM, basicPureToken)
	log.Println("@RESULT_AFTER_APPEND: ", redirect_address)
	return redirect_address
}