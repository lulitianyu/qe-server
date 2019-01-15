/**
*  Version:1.0
 * Copyright (C), 2019
 * @author Luobiken
 * @email:26015696@qq.com
 * @Tel/Weixin:13078095006
 * @Date  2019-01 14:10
 * Description：
*/
package operate

import (
	"github.com/go-xorm/xorm"
	"qe-server/db/connect"
)

func Read(command Command) (interface{}, error) {
	engine, err := connect.GetEngine()
	if err != nil {
		return nil, err
	}
	table := GetTable(command.Table)
	var rows *xorm.Rows
	//要返回的字段
	rows, err = engine.Cols(command.Column).Where(command.Where).OrderBy(command.Order).Limit(command.Limit[0], command.Limit[1]).Rows(table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables := make([]interface{}, 0)
	for rows.Next() {
		table := GetTable(command.Table)
		err = rows.Scan(table)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}
