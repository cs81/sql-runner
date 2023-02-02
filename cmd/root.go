/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sql-runner",
	Short: "sql定时执行脚本",
	Long:  `一个简单的sql定时执行脚本，需要设置db连接信息、执行间隔、运行sql`,
}

// Execute
//
//	@Description: 执行
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
