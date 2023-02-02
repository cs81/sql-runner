package api

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlInfo struct {
	SqlInfo
}

// Run
//
//	@Description: 运行mysql
//	@receiver receiver
func (receiver *MysqlInfo) Run() {
	url := receiver.GetDriver()
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		os.Exit(2)
	}
	receiver.DoRun(db)
}

// GetDriver
//
//	@Description: 获取mysql驱动
//	@receiver receiver
//	@return string
func (receiver *MysqlInfo) GetDriver() string {
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", receiver.User, receiver.Password, receiver.Host, receiver.Port, receiver.DbName)
	return url
}

// init
//
//	@Description: 初始化mysql策略
func init() {
	SqlInfoCache[MysqlDb] = func(info *SqlInfo) Runner {
		return &MysqlInfo{
			*info,
		}
	}
}
