package api

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type MysqlInfo struct {
	SqlInfo
}

func (receiver *MysqlInfo) Run() {
	url := receiver.GetDriver()
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		os.Exit(2)
	}
	receiver.DoRun(db)
}

func (receiver *MysqlInfo) GetDriver() string {
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", receiver.User, receiver.Password, receiver.Host, receiver.Port, receiver.DbName)
	return url
}

func init() {
	SqlInfoCache[MysqlDb] = func(info *SqlInfo) Runner {
		return &MysqlInfo{
			*info,
		}
	}
}
