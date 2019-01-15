/**
*  Version:1.0
 * Copyright (C), 2019
 * @author Luobiken
 * @email:26015696@qq.com
 * @Tel/Weixin:13078095006
 * @Date  2019-01 15:18
 * Description：
*/
package main

import (
	"qe-server/db/operate"
	"qe-server/service"
)

func main() {

	//同步更新表结构
	operate.SyncTable()
	/*
	   	var test = `[{
	   		"uuid":"fasdfaf243--",
	   		"password":"sdfs222df-",
	   		"account":82342342423434,
	   		"name":"asd4234fa--",
	   		"phone":"fsd4234234fs--",
	   		"email":"fsdf42342ds--",
	   		"company_name":"fsd4234234fdsf--"
	   },{
	   		"uuid":"fasdfaf",
	   		"password":"sdfsdf",
	   		"account":8234234234,
	   		"name":"asdfa",
	   		"phone":"fsdf111s",
	   		"email":"fsdfds",
	   		"company_name":"fsdfdsf"
	   }]`

	   	log.Println("开始运行了")
	   	var curd operate.Command
	   	curd.Method = "c"
	   	curd.Table = "User"
	   	curd.Data = test
	   	curd.Where = "account=0"
	   	rows, err := operate.Update(curd)
	   	if len(err) > 0 {
	   		for i, v := range err {
	   			log.Println(i, v)
	   		}
	   	}
	   	log.Println("成功插入数据:", rows)

	   	curd = *new(operate.Command)
	   	curd.Method = "r"
	   	curd.Table = "User"
	   	curd.Where = "id>0"
	   	curd.Order = "id desc"
	   	//curd.Column = "account,id"
	   	result, err2 := operate.Read(curd)
	   	if err2 != nil {
	   		log.Println(err2.Error())
	   	} else {
	   		r, _ := json.Marshal(result)
	   		log.Println("查询结果:", string(r))
	   	}
	   	curd.Method = "d"
	   	curd.Where = "account =2342342423434"
	   	affected, err2 := operate.Delete(curd)
	   	if err2 != nil {
	   		log.Println(err2.Error())
	   	} else {
	   		log.Println("删除结果:", string(affected))
	   	}
	*/
	service.WebsocketRun()

	//document.Parse()
}
