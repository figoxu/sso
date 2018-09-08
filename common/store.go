package common

import (
	"github.com/gin-contrib/sessions"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

const (
	SSO_BASIC_PURE_TOKEN = "basic_raw_token"
	SSO_TOKEN_COOKIE     = "sso"
	SSO_TOKEN_PARAM = "basic_pure_token"
	SSO_FROM_PARAM = "from"
)

func WriteCookie(c *gin.Context, basicPureToken, domain string) {
	cookie := &http.Cookie{
		Name:   SSO_TOKEN_COOKIE,
		Value:  basicPureToken,
		Path:   "/",
		Domain: domain,
		MaxAge: 60 * 60 * 24,
	}
	http.SetCookie(c.Writer, cookie)
}

func StoreToken2Session(c *gin.Context, basicPureToken string) {
	session := sessions.Default(c)
	session.Set(SSO_BASIC_PURE_TOKEN, basicPureToken)
	session.Save()
}

func GetBasicPureToken(c *gin.Context) (basicPureToken string) {
	getTokenByURL := func()string{
		vs := c.Request.URL.Query()
		return vs.Get(SSO_TOKEN_PARAM);
	}
	getTokenByCookie:= func()string{
		if v,err:=c.Cookie(SSO_TOKEN_COOKIE);err==nil{
			return v
		}
		return ""
	}
	getTokenBySession:= func() string{
		session := sessions.Default(c)
		s_token:=session.Get(SSO_BASIC_PURE_TOKEN)
		if s_token!=nil {
			return fmt.Sprint(s_token)
		}
		return ""
	}
	basicPureToken = getTokenByURL()
	if basicPureToken=="" {
		basicPureToken = getTokenByCookie()
	}
	if basicPureToken==""{
		basicPureToken = getTokenBySession()
	}
	return basicPureToken
}
