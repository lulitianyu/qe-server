/**
*  Version:1.0
 * Copyright (C), 2019
 * @author Luobiken
 * @email:26015696@qq.com
 * @Tel/Weixin:13078095006
 * @Date  2019-01 8:40
 * Description：同步struct到数据库表，跟据表名称获取对应的struct实例
*/
package service

import (
	"log"
	"qe-server/db/operate"
	"qe-server/table"
)

func SyncTable() {
	var user table.User
	err := operate.CreateTable(&user)
	if err != nil {
		log.Println(`同步表:"User"出错了`, err.Error())
	}
}
func GetTable(name string) interface{} {
	switch name {
	case "User":
		return new(table.User)
	default:
		return nil
	}
}
