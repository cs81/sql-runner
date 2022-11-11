package api

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type PgsqlInfo struct {
	SqlInfo
}

func (receiver *PgsqlInfo) Run() {
	url := receiver.GetDriver()
	db, err := sql.Open("postgres", url)
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

func (receiver *PgsqlInfo) GetDriver() string {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", receiver.Host, receiver.Port, receiver.User, receiver.Password, receiver.DbName)
	return url
}

func init() {
	SqlInfoCache[PgsqlDb] = func(info *SqlInfo) Runner {
		return &PgsqlInfo{
			*info,
		}
	}
}
