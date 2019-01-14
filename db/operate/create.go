/**
*  Version:1.0
 * Copyright (C), 2019
 * @author Luobiken
 * @email:26015696@qq.com
 * @Tel/Weixin:13078095006
 * @Date  2019-01 14:52
 * Description：插入一条新的记录到数据表
*/
package operate

import (
	"encoding/json"
	"fmt"
	"github.com/antonholmquist/jason"
	"qe-server/db/connect"
	"qe-server/service"
)

func Create(command service.Command) (int64, []error) {
	errList := make([]error, 0)
	jsonData, err := jason.NewValueFromBytes([]byte(command.Data))
	if err != nil {
		return 0, append(errList, fmt.Errorf("数据不是有效的json字符串!%s", err.Error()))
	}
	//检测是否为数组json
	jsonArray, err := jsonData.Array()
	if err != nil { //不是数组，就自己造一个数组
		var jsonArray [1]*jason.Value
		s, _ := jason.NewValueFromBytes([]byte(command.Data))
		jsonArray[0] = s
	}
	return insertMultiple(command, jsonArray)
}
func insertMultiple(command service.Command, jsonArray []*jason.Value) (int64, []error) {
	errList := make([]error, 0)
	engine, err := connect.GetEngine()
	if err != nil {
		return 0, append(errList, err)
	}
	var affected int64 = 0
	for i := 0; i < len(jsonArray); i++ {
		s, err := jsonArray[i].Object()
		if err != nil {
			//log.Println("写入数据出错了:",err.Error())
			errList = append(errList, err)
			continue
		}
		table := service.GetTable(command.Table)
		err = json.Unmarshal([]byte(s.String()), &table)
		if err != nil {
			//log.Println("出错了,单行数据不是有效的json:",err.Error())
			errList = append(errList, err)
		} else {
			row, err := engine.Insert(table)
			if err != nil {
				//log.Println("写入数据出错了:",err.Error())
				errList = append(errList, err)
			} else {
				//log.Println("成功插入数据:",row)
				affected = affected + row
			}
		}
	}
	return affected, errList

}
