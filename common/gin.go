package common

import (
	"github.com/quexer/utee"
	"github.com/figoxu/sso/pb/sso"
	"github.com/figoxu/Figo"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/url"
	"net/http"
	"google.golang.org/grpc"
	"context"
)

func BuildMid(domain, sso_redirect_url string, grpc_client *grpc.ClientConn, token_cache *utee.TimerCache) (func(c *gin.Context)) {
	saveTokenUser := func(basicPureToken string, user *sso.LoginInfoRsp) {
		token_cache.Put(basicPureToken, user)
	}

	getTokenUser := func(basicPureToken string) (user *sso.LoginInfoRsp) {
		v := token_cache.Get(basicPureToken)
		if v == nil {
			return nil
		}
		user = v.(*sso.LoginInfoRsp)
		return user
	}

	midSSo := func(c *gin.Context) {
		urlEncode := func() string {
			querystring := c.Request.URL.String()
			querystring = fmt.Sprint("http://", domain, querystring)
			return url.QueryEscape(querystring)
		}
		fmt.Println(urlEncode())

		autoSaveToken := func() {
			vs := c.Request.URL.Query()
			if basicPureToken := vs.Get(SSO_TOKEN_PARAM); basicPureToken != "" {
				WriteCookie(c, basicPureToken, domain)
				StoreToken2Session(c, basicPureToken)
			}
		}

		authCode := func() bool {
			basic_pure_token := GetBasicPureToken(c)
			if basic_pure_token == "" {
				return false
			}
			if u := getTokenUser(basic_pure_token); u != nil && u.User.Id > 0 {
				return true
			}

			ssoServiceClient := sso.NewSsoServiceClient(grpc_client)
			req := &sso.LoginInfoReq{
				BasicRawToken: basic_pure_token,
			}
			rsp, err := ssoServiceClient.GetLoginInfo(context.Background(), req)
			if err != nil || rsp.User.Id < 0 {
				return false
			}
			saveTokenUser(basic_pure_token, rsp)
			return true
		}

		autoSaveToken()
		if !authCode() {
			rurl := Figo.UrlAppendParam(sso_redirect_url, SSO_FROM_PARAM, urlEncode())
			c.Redirect(http.StatusFound, rurl)
			return
		} else {
			c.Next()
		}
	}
	return midSSo
}
