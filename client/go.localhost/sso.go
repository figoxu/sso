package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/url"
	"github.com/figoxu/sso/common"
	"github.com/figoxu/sso/pb/sso"
	"context"
	"github.com/figoxu/Figo"
	"net/http"
)

func midSSo(c *gin.Context) {
	urlEncode := func() string {
		querystring := c.Request.URL.String()
		querystring = fmt.Sprint("http://", sysEnv.domain, querystring)
		return url.QueryEscape(querystring)
	}
	fmt.Println(urlEncode())

	autoSaveToken := func() {
		vs := c.Request.URL.Query()
		if basicPureToken := vs.Get(common.SSO_TOKEN_PARAM); basicPureToken != "" {
			common.WriteCookie(c, basicPureToken, sysEnv.domain)
			common.StoreToken2Session(c, basicPureToken)
		}
	}

	authCode := func() bool {
		basic_pure_token := common.GetBasicPureToken(c)
		if basic_pure_token == "" {
			return false
		}
		if u := getTokenUser(basic_pure_token); u != nil && u.User.Id > 0 {
			return true
		}

		ssoServiceClient := sso.NewSsoServiceClient(sysEnv.grpc_client)
		req := &sso.LoginInfoReq{
			BasicRawToken: basic_pure_token,
		}
		rsp,err:= ssoServiceClient.GetLoginInfo(context.Background(), req)
		if err!=nil || rsp.User.Id < 0 {
			return false
		}
		saveTokenUser(basic_pure_token,rsp)
		return true
	}

	autoSaveToken()
	if !authCode() {
		rurl:=Figo.UrlAppendParam(sysEnv.sso_redirect_url,common.SSO_FROM_PARAM,urlEncode())
		c.Redirect(http.StatusFound,rurl)
		return
	} else {
		c.Next()
	}
}

func saveTokenUser(basicPureToken string, user *sso.LoginInfoRsp) {
	sysEnv.token_cache.Put(basicPureToken, user)
}

func getTokenUser(basicPureToken string) (user *sso.LoginInfoRsp) {
	v := sysEnv.token_cache.Get(basicPureToken)
	user = v.(*sso.LoginInfoRsp)
	return user
}


