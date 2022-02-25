// @Author: su
// @Date: 2022-2-21 , 0021

package models

import "time"

type User struct {
	Id 		int64
	Name 	string 		`xorm:"varchar(25) notnull unique 'user_name' comment(姓名)"`
	Addr 	string 		`xorm:"varchar(150) comment(地址)"`
	Phone 	string 		`xorm:"varchar(11) comment(手机号)"`
	GroupId int64 		`xorm:"index comment(组名)"`
	TypeId 	int64 		`xorm:"index comment(类型名)"`
	CreatTime 	time.Time 	`xorm:"created comment(创建时间)"`
	UpdateTime 	time.Time 	`xorm:"updated comment(更新时间)"`
	DeleteTime 	time.Time 	`xorm:"deleted comment(删除时间)"`
}
