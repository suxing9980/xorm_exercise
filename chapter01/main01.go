// @Author: su
// @Date: 2022-2-21 , 0021

package main

import (
	"fmt"
	"time"
	"xorm.io/xorm"
	"xorm_exercise/database"
	"xorm_exercise/models"
)

var engine *xorm.Engine

func main() {
	engine = database.InitMysql()
	defer engine.Close()

	// 插入1条数据
	user01 := models.User{
		Name:       "wangliang",
		Addr:       "福田路一号",
		Phone:      "13695284671",
		CreatTime:  time.Now(),
		UpdateTime: time.Now(),
		DeleteTime: time.Now(),
	}
	one, err := engine.InsertOne(&user01)
	if err != nil {
		fmt.Println("插入一条数据ID错误:",err)
	}
	fmt.Println("插入一条数据ID:",one)

	//user02 := models.User{
	//	Name:       "liming",
	//	Addr:       "福田路二号",
	//	Phone:      "13695284570",
	//	CreatTime:  time.Now(),
	//	UpdateTime: time.Now(),
	//	DeleteTime: time.Now(),
	//}
	//in, err := engine.Insert(&user02)
	//if err != nil {
	//	fmt.Println("插入一条数据ID错误:",err)
	//}
	//fmt.Println("插入一条数据ID:",in)
}
