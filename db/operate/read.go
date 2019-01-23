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
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"qe-server/db/connect"
)

func Read(command Command) (interface{}, error) {
	engine, err := connect.GetEngine()
	if err != nil {
		return nil, err
	}
	table := GetTable(command.Table)
	if table == nil {
		return nil, errors.New(`无效的数据表名称:"` + command.Table + `"`)
	}
	var rows *xorm.Rows
	//要返回的字段
	fmt.Println("table", table)
	var session *xorm.Session
	if command.Column != "" {
		session = engine.Cols(command.Column)
	}
	if command.Where != "" {
		session = session.Where(command.Where)
	}
	if command.Order != "" {
		session = session.OrderBy(command.Order)
	}
	if command.Limit[0] != 0 && command.Limit[1] != 0 {
		session = session.Limit(command.Limit[0], command.Limit[1])
	}
	if session == nil {
		rows, err = engine.Rows(table)
	} else {
		rows, err = session.Rows(table)
	}
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
