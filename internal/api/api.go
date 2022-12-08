package api

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

const (
	MysqlDb = "mysql"
	PgsqlDb = "pgsql"
)

var (
	SqlInfoCache = map[string]func(info *SqlInfo) Runner{}
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
	Cron     string
}

func (receiver *SqlInfo) DoRun(db *sql.DB) {
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	if receiver.Cron != "" {
		c := cron.New(cron.WithLocation(time.UTC))
		_, err := c.AddFunc(receiver.Cron, func() {
			runSql(db, receiver.Sql)
		})
		defer c.Stop()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		c.Start()
	} else {
		duration := time.Second * time.Duration(receiver.Second)
		t := time.NewTimer(duration)
		defer t.Stop()
		go func() {
			for {
				<-t.C
				runSql(db, receiver.Sql)
				// 需要重置Reset 使 t 重新开始计时
				t.Reset(duration)
			}
		}()
	}
	select {}
}

func runSql(db *sql.DB, sql string) {
	fmt.Printf("%v => 运行sql：%v\n", time.Now(), sql)
	query, err := db.Query(sql)
	if err != nil {
		fmt.Printf("%v => %v\n", time.Now(), err.Error())
		os.Exit(2)
	}
	fmt.Printf("%v => %v\n", time.Now(), "运行结果成功")
	_ = query.Close()
}
