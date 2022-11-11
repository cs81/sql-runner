package api

import (
	"database/sql"
	"fmt"
	"time"
)

type Runner interface {
	Run()
	GetDriver() string
}

type SqlInfo struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	Sql      string
	Second   int
	DbType   string
}

func (receiver *SqlInfo) DoRun(db *sql.DB) {
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

const (
	MysqlDb = "mysql"
	PgsqlDb = "pgsql"
)

var (
	SqlInfoCache = map[string]func(info *SqlInfo) Runner{}
)
