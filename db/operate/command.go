/**
*  Version:1.0
 * Copyright (C), 2019
 * @author Luobiken
 * @email:26015696@qq.com
 * @Tel/Weixin:13078095006
 * @Date  2019-01 14:05
 * Description：接收客户端发来的操作命令
*/
package operate

/**
插入一条新数据
{Method:"c",table:"my_table",data:{field1:"data1",field2:"data2",...} }
更新一条数据
{Method:"u",table:"my_table",data:{field1:"data1",field2:"data2",...} }
删除一条数据，其中data里面的字段就是删除条件
{Method:"d",table:"my_table",data:{field1:"data1",field2:"data2",...} }
查询，其中的data就是查询条件
{Method:"r",table:[{table1:["field1","field2",...]},{table2:*},...],data:{field1:"data1",field2:"data2",...} }
*/
type Command struct {
	Method string //指示要进行什么操作，创建c,更新u，读取r,删除d
	Table  string //要操作的表
	Column string //查询要返回的字段
	Data   string //操作的数据
	Where  string //操作条件
	Order  string //排序方法
	Limit  [2]int //条数限制，如[100,10]，从第10条起，最多100条---xorm限制获取的数目，第一个参数为条数，第二个参数表示开始位置，如果不传则为0
}

//一般输出
type Result struct {
	State   bool
	Message string
	Data    interface{}
}

//分页输出
type Page struct {
	RowCount  int64
	PageSize  int64
	PageCount int64
	PageIndex int64
	DataRow   interface{}
}
