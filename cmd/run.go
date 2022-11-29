package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sql-runner/internal/api"
)

var info = &api.SqlInfo{}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "运行",
	Long:  `开始运行sql`,
	Run: func(cmd *cobra.Command, args []string) {
		initFunc := api.SqlInfoCache[info.DbType]
		if initFunc == nil {
			fmt.Println("不支持的DB")
			os.Exit(2)
		}
		initFunc(info).Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVarP(&info.Second, "second", "S", 10, "多久执行一次，单位秒，多少秒执行一次sql")
	runCmd.Flags().StringVarP(&info.Cron, "cron", "C", "", "cron表达式，按表达式执行sql")
	runCmd.Flags().StringVarP(&info.Host, "host", "H", "127.0.0.1", "连接主机")
	runCmd.Flags().IntVarP(&info.Port, "port", "P", 3306, "连接端口")
	runCmd.Flags().StringVarP(&info.DbType, "dbType", "D", "mysql", "数据库类型:mysql、pgsql")
	runCmd.Flags().StringVarP(&info.User, "user", "u", "root", "连接用户名")
	runCmd.Flags().StringVarP(&info.Password, "password", "p", "root", "连接密码")
	runCmd.Flags().StringVarP(&info.DbName, "dbName", "d", "", "连接数据库")
	_ = runCmd.MarkFlagRequired("dbName")
	runCmd.Flags().StringVarP(&info.Sql, "sql", "s", "", "定时执行sql")
	_ = runCmd.MarkFlagRequired("sql")

}
