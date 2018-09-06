package main

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"reflect"
	"github.com/figoxu/sso/pb/sso"
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

func (p *Resource) toProto() *sso.Resource {
	resource := &sso.Resource{
		Id:         int32(p.Id),
		Pid:        int32(p.Pid),
		Name:       p.Name,
		SysName:    p.SysName,
		Priority:   int32(p.Priority),
		Path:       p.Path,
		Type:       p.Type,
		Permission: p.Permission,
		Available:  p.Available,
	}
	return resource
}

type ResourceDao struct {
	db *gorm.DB
}

func NewResourceDao(db *gorm.DB) *ResourceDao {
	return &ResourceDao{
		db: db,
	}
}

func (p *ResourceDao) FindByUid(uid int) []Resource {
	resources := make([]Resource, 0)
	return resources
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
	Gids         IntArray `sql:"type:integer[]"`
	Available    bool
}

func (p *User) toProto() *sso.User {
	return &sso.User{
		Id:        int32(p.Id),
		Name:      p.Name,
		Phone:     p.Phone,
		Gids:      convertIntArray2Int32Array(p.Gids),
		Available: p.Available,
	}
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

func (p *UserDao) Save(user *User)  {
	p.db.Save(&user)
}

type UserGroup struct {
	Id        int
	Name      string
	Resources IntArray `sql:"type:integer[]"`
	Available bool
}

func (p *UserGroup) toProto() *sso.UserGroup {
	ug := &sso.UserGroup{
		Id:        int32(p.Id),
		Name:      p.Name,
		Resources: convertIntArray2Int32Array(p.Resources),
		Available: p.Available,
	}
	return ug
}

type UserGroupDao struct {
	db *gorm.DB
}

func NewUserGroupDao(db *gorm.DB) *UserGroupDao {
	return &UserGroupDao{
		db: db,
	}
}

func (p *UserGroupDao) FindByUid(uid int) []UserGroup {
	ugs := make([]UserGroup, 0)
	return ugs
}
