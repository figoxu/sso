package main

import (
	"github.com/gin-gonic/gin"
	"encoding/hex"
	"github.com/pborman/uuid"
	"github.com/figoxu/Figo"
	"net/http"
	"github.com/quexer/utee"
	"fmt"
	"strings"
	"github.com/gin-contrib/sessions"
	"encoding/base64"
)

const (
	SSO_FROM            = "sso_from"
	SSO_BASIC_RAW_TOKEN = "basic_raw_token"
	SSO_TOKEN_COOKIE    = "sso"
)

type TokenHelper struct {
	ctx     *gin.Context
	userDao *UserDao
	domain  string
}

func NewTokenHelper(c *gin.Context) TokenHelper {
	return TokenHelper{
		ctx:     c,
		userDao: NewUserDao(pg_rbac),
	}
}

func (p *TokenHelper) NewToken(uid int) string {
	genToken := func() (basicRawToken string) {
		user := p.userDao.GetById(uid)
		user.Token = uuid.New()
		user.TokenSalt = uuid.New()
		p.userDao.Update(user, Figo.SnakeStrings("Token", "TokenSalt")...)
		tokenSaltHelper := NewUserTokenSaltHelper(user)
		rawToken := tokenSaltHelper.Encode(user.Token)
		return rawToken
	}
	writeCookie := func(basicRawToken string) {
		cookie := &http.Cookie{
			Name:   SSO_TOKEN_COOKIE,
			Value:  basicRawToken,
			Path:   "/",
			Domain: sysEnv.domain,
			MaxAge: 60 * 60 * 24,
		}
		http.SetCookie(p.ctx.Writer, cookie)
	}
	storeToken2Session := func(basicRawToken string) {
		session := sessions.Default(p.ctx)
		session.Set(SSO_BASIC_RAW_TOKEN, basicRawToken)
	}
	rawToken := genToken()
	basicRawToken := BasicAuthEncode(NewBasicAuthStr(uid, rawToken))
	writeCookie(basicRawToken)
	storeToken2Session(basicRawToken)
	return rawToken
}

func ParseToken(basicRawToken string) (uid int, rawToken string) {
	content := BasicAuthDecode(basicRawToken)
	uid, rawToken = ParseBasicAuth(content)
	return uid, rawToken
}

func CheckRawToken(uid int, rawToken string) bool {
	userDao:=NewUserDao(pg_rbac)
	user := userDao.GetById(uid)
	tokenSaltHelper := NewUserTokenSaltHelper(user)
	token := tokenSaltHelper.Encode(rawToken)
	return user.Token == token
}

type SaltHelper struct {
	basic string
	plus  string
}

func NewSaltHelper(basic, plus string) SaltHelper {
	return SaltHelper{
		basic: basic,
		plus:  plus,
	}
}

func NewUserTokenSaltHelper(user User) SaltHelper {
	return NewSaltHelper(fmt.Sprint(user.Id, user.TokenSalt), user.TokenSalt)
}

func NewUserPasswordSaltHelper(user User) SaltHelper {
	return NewSaltHelper(fmt.Sprint(user.Id, user.PasswordSalt), user.PasswordSalt)
}

func (p *SaltHelper) Encode(content string) (result string) {
	encry := func(s string) string {
		aesHelper := Figo.NewAesHelp(utee.Md5([]byte(p.basic)), utee.Md5([]byte(p.plus))...)
		bs := []byte(s)
		result, err := aesHelper.Encrypt(bs)
		utee.Chk(err)
		return hex.EncodeToString(result)
	}
	return encry(content)
}

func (p *SaltHelper) Decode(content string) (result string) {
	bs, err := hex.DecodeString(content)
	utee.Chk(err)
	aesHelper := Figo.NewAesHelp(utee.Md5([]byte(p.basic)), utee.Md5([]byte(p.plus))...)
	bs, err = aesHelper.Decrypt(bs)
	utee.Chk(err)
	return string(bs)
}

func NewBasicAuthStr(id int, token string) string {
	return fmt.Sprint(id, ":", token)
}

func ParseBasicAuth(auth string) (id int, token string) {
	ary := strings.Split(auth, ":")
	if len(ary) != 2 {
		err := fmt.Errorf("bad auth string %s\n", auth)
		utee.Chk(err)
	}
	id, err := Figo.TpInt(ary[0])
	utee.Chk(err)
	return id, token
}

func BasicAuthEncode(content string) string {
	return base64.StdEncoding.EncodeToString([]byte(content))
}

func BasicAuthDecode(content string) string {
	bs, err := base64.StdEncoding.DecodeString(content)
	utee.Chk(err)
	return string(bs)
}
