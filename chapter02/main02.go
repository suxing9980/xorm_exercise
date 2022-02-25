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

	// 1 ID get
	user01 := new(models.User)
	u01, _ := engine.ID(1).Get(user01)
	fmt.Println(u01,user01)

	// 2  where get
	user02 := new(models.User)
	u02, _ := engine.Where("addr = ?", "福田路一号").Get(user02)
	fmt.Println(u02, user02)

	// 3 根据结构体中的非空字段来获取
	user03 := models.User{Id: 2}
	u03, _ := engine.Get(&user03)
	fmt.Println(u03, user03)

	// 4 Find slice
	users01 := make([]models.User, 0)
	err := engine.Where("Id < ?", 3).Find(&users01)
	if err != nil {
		fmt.Println("Find error",err)
	}
	fmt.Println(users01)

	// 5 Find map
	users02 := make(map[int64]models.User)
	err = engine.Where("Id < ?", 3).Find(&users02)
	if err != nil {
		fmt.Println("Find error",err)
	}
	fmt.Println("4:", users02)

	// 6 针对单个的字段
	var names []string
	engine.Table("user").Cols("user_name").Find(&names)
	fmt.Println("6", names)
}
