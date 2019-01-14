/**
*  Version:1.0
 * Copyright (C), 2019
 * @author Luobiken
 * @email:26015696@qq.com
 * @Tel/Weixin:13078095006
 * @Date  2019-01 14:58
 * Description：从struct创建数据库表的方法
*/
package operate

import "qe-server/db/connect"

func CreateTable(beans interface{}) error {
	engine, err := connect.GetEngine()
	if err != nil {
		return err
	}
	//同步一下数据表
	return engine.Sync2(beans)
}
