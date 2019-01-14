/**
*  Version:1.0
 * Copyright (C), 2019
 * @author Luobiken
 * @email:26015696@qq.com
 * @Tel/Weixin:13078095006
 * @Date  2019-01 9:45
 * Description：
*/
package table

//要注意，struct体的字段开头一定要大写，不映射到数据表会出错！！bug
type User struct {
	Id          int64  `xorm:"pk autoincr"`
	Uuid        string `xorm:"varchar(256) default('')"`
	Password    string `xorm:"varchar(256) default('')"`
	Account     int64  `xorm:"unique default(0)"`
	Name        string `xorm:"varchar(64) default('')"`
	Phone       string `xorm:"unique varchar(16) default('')"`
	Email       string `xorm:"unique varchar(128) default('')"`
	CompanyName string `xorm:"varchar(128) default('')"`
}
