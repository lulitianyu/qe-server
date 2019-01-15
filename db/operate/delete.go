/**
*  Version:1.0
 * Copyright (C), 2019
 * @author Luobiken
 * @email:26015696@qq.com
 * @Tel/Weixin:13078095006
 * @Date  2019-01 15:22
 * Description：
*/
package operate

import (
	"qe-server/db/connect"
)

func Delete(command Command) (int64, error) {
	engine, err := connect.GetEngine()
	if err != nil {
		return 0, err
	}
	table := GetTable(command.Table)
	//要返回的字段
	return engine.Where(command.Where).Delete(table)
}
