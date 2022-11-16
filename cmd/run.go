package cmd

import (
	"github.com/spf13/cobra"
	"sql-runner/api"
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
			panic("不支持的DB")
		}
		initFunc(info).Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVarP(&info.Second, "second", "S", 10, "多久执行一次，单位秒")
	runCmd.Flags().StringVarP(&info.Host, "host", "H", "127.0.0.1", "连接主机")
	_ = runCmd.MarkFlagRequired("host")
	runCmd.Flags().IntVarP(&info.Port, "port", "P", 3306, "连接端口")
	_ = runCmd.MarkFlagRequired("port")
	runCmd.Flags().StringVarP(&info.DbType, "dbType", "D", "mysql", "数据库类型:mysql、pgsql")
	runCmd.Flags().StringVarP(&info.User, "user", "u", "root", "连接用户名")
	_ = runCmd.MarkFlagRequired("user")
	runCmd.Flags().StringVarP(&info.Password, "password", "p", "root", "连接密码")
	_ = runCmd.MarkFlagRequired("password")
	runCmd.Flags().StringVarP(&info.DbName, "dbName", "d", "", "连接数据库")
	_ = runCmd.MarkFlagRequired("dbName")
	runCmd.Flags().StringVarP(&info.Sql, "sql", "s", "", "定时执行sql")
	_ = runCmd.MarkFlagRequired("sql")

}
