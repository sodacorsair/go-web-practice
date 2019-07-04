package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Userinfo struct {
	Uid        int `orm:"PK"`
	Username   string
	Departname string
	Created    time.Time
}

type User struct {
	Uid     int `orm:"PK"`
	Name    string
	Profile *Profile `orm:"rel(one)"`
	Post    []*Post  `orm:"reverse(many)"`
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"`
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		"root:225500@/goweb?charset=utf8", 30)
	orm.RegisterModel(new(Userinfo), new(User), new(Profile), new(Post), new(Tag))
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	users := []User{
		{Name: "slene"},
		{Name: "astaxie"},
		{Name: "unknown"},
	}

	successNum, err := o.InsertMulti(3, users)
	fmt.Println(successNum)
	fmt.Println(err)
}
