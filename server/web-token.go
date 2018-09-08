package main

import (
	"github.com/gin-gonic/gin"
	"encoding/hex"
	"github.com/pborman/uuid"
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"fmt"
	"strings"
	"encoding/base64"
)

const (
	SSO_FROM             = "sso_from"
	SSO_BASIC_PURE_TOKEN = "basic_raw_token"
	SSO_TOKEN_COOKIE     = "sso"
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
		pureToken := uuid.New()
		user := p.userDao.GetById(uid)
		user.TokenSalt = uuid.New()
		tokenSaltHelper := NewUserTokenSaltHelper(user)
		user.Token = tokenSaltHelper.Encode(pureToken)
		p.userDao.Update(user, Figo.SnakeStrings("Token", "TokenSalt")...)
		return pureToken
	}
	pureToken := genToken()
	basicPureToken := BasicAuthEncode(NewBasicAuthStr(uid, pureToken))
	return basicPureToken
}

func ParseToken(basicRawToken string) (uid int, pureToken string) {
	content := BasicAuthDecode(basicRawToken)
	uid, pureToken = ParseBasicAuth(content)
	return uid, pureToken
}

func CheckPureToken(uid int, rawToken string) bool {
	userDao := NewUserDao(pg_rbac)
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

func NewUserTokenSaltHelper(user UserAccount) SaltHelper {
	return NewSaltHelper(fmt.Sprint(user.Id, user.TokenSalt), user.TokenSalt)
}

func NewUserPasswordSaltHelper(user UserAccount) SaltHelper {
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
