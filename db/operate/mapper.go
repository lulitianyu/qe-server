/**
*  Version:1.0
 * Copyright (C), 2019
 * @author Luobiken
 * @email:26015696@qq.com
 * @Tel/Weixin:13078095006
 * @Date  2019-01 8:40
 * Description：同步struct到数据库表，跟据表名称获取对应的struct实例
*/
package operate

import (
	"log"
	"qe-server/table"
)

func SyncTable() {
	var user table.User
	err := CreateTable(&user)
	if err != nil {
		log.Println(`同步表:"User"出错了`, err.Error())
	}

	var project table.Project
	err = CreateTable(&project)
	if err != nil {
		log.Println(`同步表:"ArticleSort"出错了`, err.Error())
	}

	var article table.Article
	err = CreateTable(&article)
	if err != nil {
		log.Println(`同步表:"Article"出错了`, err.Error())
	}
}
func GetTable(name string) interface{} {
	switch name {
	case "User":
		return new(table.User)
	case "Project":
		return new(table.Project)
	case "Article":
		return new(table.Article)
	default:
		return nil
	}
}
