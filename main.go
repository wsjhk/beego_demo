package main

import (
	"fmt"

	"../beego_project/models"
	_ "../beego_project/routers" //import时候带有"_"下划线表示只执行该包的所有init()函数，无法通过报名.函数来引用函数和变量
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	maxIdle := 30
	maxConn := 30
	dbhost := beego.AppConfig.String("mysqlhost")
	dbport := beego.AppConfig.String("mysqlport")
	dbuser := beego.AppConfig.String("mysqluser")
	dbpassword := beego.AppConfig.String("mysqlpass")
	db := beego.AppConfig.String("mysqldb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn, maxIdle, maxConn)
	fmt.Printf("数据库连接成功！%s\n", conn)
	//如果第二个参数是false，那么表存在时就会skip；如果为true，就是删除重建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	profile := new(models.Profile)
	profile.Age = 30

	user := new(models.User)
	user.Profile = profile
	user.Name = "jhk"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
	beego.Run()

}
