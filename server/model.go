package main

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"reflect"
)

var (
	pg_rbac *gorm.DB
)

func initDb(dbUrl string) {
	db_rbac, err := gorm.Open("postgres", dbUrl)
	db_rbac.DB().SetConnMaxLifetime(time.Minute * 5)
	db_rbac.DB().SetMaxIdleConns(0)
	db_rbac.DB().SetMaxOpenConns(5)
	db_rbac.SetLogger(&Figo.GormLog{})
	utee.Chk(err)
	db_rbac.LogMode(true)
	db_rbac.SingularTable(true)
	db_rbac.Debug().AutoMigrate(Entities()...)
	pg_rbac = db_rbac
}

func Entities() (values []interface{}) {
	values = append(values, &Resource{})
	values = append(values, &User{})
	values = append(values, &UserGroup{})
	return values
}

type Resource struct {
	Id         int
	Pid        int
	Name       string
	SysName    string
	Priority   int
	Path       string
	Type       string //菜单、按钮
	Permission string
	Available  bool
}

type User struct {
	Id           int
	Account      string
	Password     string
	PasswordSalt string
	Token        string
	TokenSalt    string
	Name         string
	Phone        string
	Gids         []int
	Available    bool
}

type UserGroup struct {
	Id        int
	Name      string
	Resources []int
	Available bool
}

type ResourceDao struct {
	db *gorm.DB
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (p *UserDao) GetByLoginName(loginName string) (user User) {
	sb := Figo.NewSqlBuffer()
	sb.Append(" SELECT * FROM user WHERE name=? ", loginName)
	sb.Append(" OR phone=? ", loginName)
	p.db.Raw(sb.SQL(), sb.Params()...).Scan(&user)
	return user
}

func (p *UserDao) GetById(id int) (user User) {
	p.db.Raw("SELECT * FROM user WHERE id=?", id).Scan(&user)
	return user
}

func (p *UserDao) Update(user User, fields ...string) {
	if recordNil := user.Id <= 0; recordNil {
		return
	}
	immutable := reflect.ValueOf(user)
	dataMap := make(map[string]interface{})
	for _, field := range fields {
		prop := Figo.CamelString(field)
		dataMap[field] = immutable.FieldByName(prop).Interface()
	}
	p.db.Model(&user).Select(fields).Update(dataMap)
}

type UserGroupDao struct {
	db *gorm.DB
}
