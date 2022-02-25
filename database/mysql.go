// @Author: su
// @Date: 2022-2-21 , 0021

package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
	"xorm_exercise/models"
)

var engine *xorm.Engine

func InitMysql() *xorm.Engine{
	addr := "admin:123456@tcp(192.168.124.210:3306)/xorm?charset=utf8mb4"
	engine, err := xorm.NewEngine("mysql", addr)
	if err != nil {
		fmt.Println(err.Error())
	}

	cacher := caches.NewLRUCacher(caches.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	err = engine.Ping()
	if err != nil {
		fmt.Println("连接数据库成功！")
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(log.LOG_DEBUG )	// 此处文档写engine.Logger().SetLevel(core.LOG_DEBUG)是错误的

	f, err := os.OpenFile("./log/sql.log", os.O_RDWR |os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		println(err.Error())
	}
	engine.SetLogger(log.NewSimpleLogger(f))	// 此处与文档对应不上，与log有关的需要在log包下
	engine.SetMapper(names.GonicMapper{})		// 命名规则，类似驼峰命名，ID会被翻译成id

	// 可以单独对表明与列名进行设置
	//engine.SetTableMapper(names.SameMapper{})
	//engine.SetColumnMapper(names.SnakeMapper{})

	engine.Sync2(
		new(models.User),
		new(models.Group),
		new(models.Type),
		new(models.TypeUserGroup))

	return engine
}
