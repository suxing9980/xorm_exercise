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

	user := models.User{
		Name: "LiMing002",
	}
	// 1 一般的更新
	engine.Where("user_name = ?", "liming").Update(&user)

	// 2 更新数值为0
	engine.Table(new(models.User)).Where("addr = ?", "福田路一号").Update(map[string]interface{}{"addr":"上海路18号"})
	engine.Table(new(models.User)).Where("type_id = ?", 1).Update(map[string]interface{}{"type_id":0})
	engine.Table(new(models.User)).Where("type_id = ?", 2).Cols("type_id").Update(&models.User{
		TypeId: 0,
	})

	// count
	m := new(models.Group)
	count, _ := engine.Where("id >= ?", 1).Count(m)
	fmt.Println("count:",count)

	// like
	u := models.User{}
	res, _ := engine.Where("user_name like ?", "Li"+"%").Get(&u)
	fmt.Println(res,u)
}