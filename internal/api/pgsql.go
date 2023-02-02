package api

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PgsqlInfo struct {
	SqlInfo
}

// Run
//
//	@Description: 运行pgsql
//	@receiver receiver
func (receiver *PgsqlInfo) Run() {
	url := receiver.GetDriver()
	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		os.Exit(2)
	}
	receiver.DoRun(db)
}

// GetDriver
//
//	@Description: 获取pgsql驱动
//	@receiver receiver
//	@return string
func (receiver *PgsqlInfo) GetDriver() string {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", receiver.Host, receiver.Port, receiver.User, receiver.Password, receiver.DbName)
	return url
}

// init
//
//	@Description: 初始化pgsql策略
func init() {
	SqlInfoCache[PgsqlDb] = func(info *SqlInfo) Runner {
		return &PgsqlInfo{
			*info,
		}
	}
}
