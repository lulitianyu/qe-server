package connect

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

//var dataSource = "root:Laoluo@1982@/lulit_qe?charset=utf8mb4"

//var dataSource = "root:Laoluo@1982@tcp(powerfsun.com:3306)/lulit_log?charset=utf8mb4"
var dataSource = "root:123456@/lulit_qe?charset=utf8mb4"
var engine *xorm.Engine = nil

func GetEngine() (*xorm.Engine, error) {
	//fmt.Println("run GetDBEngine...")
	if engine == nil {
		err := connect()
		if err == nil {
			return engine, nil
		} else {
			return nil, err
		}
	} else {
		err := engine.Ping()
		if err == nil {
			fmt.Println("Ping()数据库成功!")
			return engine, nil
		} else {
			err := connect()
			if err == nil {
				return engine, nil
			} else {
				return nil, err
			}
		}
	}
}
func connect() error {
	var err error
	engine, err = xorm.NewEngine("mysql", dataSource)
	if err != nil {
		return err
	} else {
		log.Println("连接数据库成功!")
		setEngine()

		return nil
	}
}
func setEngine() {
	//日志打印SQL
	engine.ShowSQL(true)
	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(100)
	//设置最大打开连接数
	engine.SetMaxOpenConns(1000)
	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	//XormEngine.SetTableMapper(core.SnakeMapper{})
}
