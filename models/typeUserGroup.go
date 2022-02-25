// @Author: su
// @Date: 2022-2-22 , 0022

package models

type TypeUserGroup struct {
	User `xorm:"extends"`
	Group `xorm:"extends"`
	Type `xorm:"extends"`
}
