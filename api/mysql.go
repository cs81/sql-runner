package api

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MysqlInfo struct {
	SqlInfo
}

func (receiver *MysqlInfo) Run() {
	url := receiver.GetDriver()
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	duration := time.Second * time.Duration(receiver.Second)
	t := time.NewTimer(duration)
	defer t.Stop()
	for {
		<-t.C
		fmt.Printf("运行sql：%v\n", receiver.Sql)
		query, err := db.Query(receiver.Sql)
		if err == nil {
			fmt.Printf("运行结果成功")
			_ = query.Close()
		}
		// 需要重置Reset 使 t 重新开始计时
		t.Reset(duration)
	}
}

func (receiver *MysqlInfo) GetDriver() string {
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", receiver.User, receiver.Password, receiver.Host, receiver.Port, receiver.DbName)
	return url
}

func init() {
	SqlInfoCache[MysqlDb] = func(info *SqlInfo) Runner {
		return &PgsqlInfo{
			*info,
		}
	}
}
