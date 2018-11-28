package models

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)
var Engine *xorm.Engine

func init(){
	user := beego.AppConfig.String("mysql_username")
	passwd := beego.AppConfig.String("mysql_pwd")
	host := beego.AppConfig.String("mysql_host")
	port,err := beego.AppConfig.Int("mysql_port")
	dbname := beego.AppConfig.String("mysql_dbname")

	if err!=nil {
		port =3306
	}
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//// set default database
	//orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
	//	user, passwd, host, port, dbname))
	//
	//orm.RegisterModel(new(Order))
	//
	//if beego.AppConfig.String("runmode") == "dev" {
	//	orm.Debug = true
	//}
	//
	//orm.RunSyncdb("default", false, false)
	Engine,_ = xorm.NewEngine("mysql",fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname))
	//if err != nil {
	//	beego.Error(err)
	//	return
	//}
	//连接测试
	if err := Engine.Ping();err != nil{
		beego.Error(err)
		return
	}

	//日志打印sql
	Engine.ShowSQL(true)

	//设置连接池的大小
	Engine.SetMaxIdleConns(5)

	//设置最大打开的连接数
	Engine.SetMaxOpenConns(5)


	//名称映射规则主要负责结构名称到表名和结构体field到表字段的名称
	Engine.SetTableMapper(core.SnakeMapper{})

}
