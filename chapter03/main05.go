// @Author: su
// @Date: 2022-2-23 , 0023

package main

import (
	"fmt"
	"xorm_exercise/database"
	"xorm_exercise/models"
)

func main() {
	engine := database.InitMysql()
	defer engine.Close()

	// 事务处理
	session := engine.NewSession()

	if err := session.Begin(); err != nil {
		fmt.Println(err.Error())
	}

	user := models.User{
		Name:    "zhaoyang",
		Addr:    "北京路28号",
		Phone:   "13962608571",
		GroupId: 3,
		TypeId:  1,
	}
	// 事务插入数据
	_, err := session.Insert(&user)
	if err != nil {
		session.Rollback()
		fmt.Println("insert error!")
	}
	user1 := models.User{
		Name:    "LiMing",
	}
	// 事务更新数据
	_, err = session.Where("user_name = ?", "LiMing002").Update(&user1)
	if err != nil {
		session.Rollback()
		fmt.Println("update error!")
	}
	// 事务删除数据
	_, err = session.Where("phone = ?", "13695284671").Delete(&models.User{})
	if err != nil {
		session.Rollback()
		fmt.Println("Delete error!")
	}
	session.Commit()
}