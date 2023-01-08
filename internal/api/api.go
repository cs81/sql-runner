package api

import (
	"database/sql"
	"encoding/json"
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
	fmt.Printf("%v => 运行sql: %v\n", time.Now().Format("2006-01-02 15:04:05"), sql)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Printf("%v => %v\n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
		os.Exit(2)
	}
	defer func() {
		_ = rows.Close()
	}()
	fmt.Printf("%v => %v\n", time.Now().Format("2006-01-02 15:04:05"), "运行成功")
	fmt.Printf("%v => %v\n", time.Now().Format("2006-01-02 15:04:05"), "结果如下")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println()
	var values []map[string][]byte
	rowsToStruct(rows, &values)
	if len(values) > 0 {
		for i, value := range values {
			fmt.Printf("%d: ", i+1)
			for k, v := range value {
				fmt.Printf("%s: %v ", k, string(v))
			}
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------------------")
}

func rowsToStruct[T any](rows *sql.Rows, container *[]T) {
	columns, _ := rows.Columns()
	cLen := len(columns)
	for rows.Next() {
		temp := make([]any, cLen)
		for i := 0; i < cLen; i++ {
			var column any
			temp[i] = &column
		}
		_ = rows.Scan(temp...)
		tempMap := make(map[string]any, cLen)
		for i, value := range temp {
			tempMap[columns[i]] = *value.(*any)
		}
		tempByte, _ := json.Marshal(tempMap)
		value := new(T)
		_ = json.Unmarshal(tempByte, value)
		*container = append(*container, *value)
	}
}
