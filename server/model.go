package main

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
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
	Id        int
	Name      string
	Phone     string
	Gids      []int
	Available bool
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

type UserGroupDao struct {
	db *gorm.DB
}
