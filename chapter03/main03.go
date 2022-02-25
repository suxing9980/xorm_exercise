// @Author: su
// @Date: 2022-2-21 , 0021

package main

import (
	"fmt"
	"xorm_exercise/database"
	"xorm_exercise/models"
)


func main() {
	engine := database.InitMysql()
	defer engine.Close()

	g := make([]models.TypeUserGroup, 0)
	err := engine.Table("user").Join("inner","group", "group.id = user.group_id").
		Join("inner", "type", "type.id = user.type_id").Find(&g)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(g)
}
