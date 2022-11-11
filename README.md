# sql-runner

#### 介绍
一个简单的sql定时执行器，支持mysql、pgsql

#### 获取方式

1. gitee地址: https://gitee.com/cailiangchen/sql-runner.git
2. github地址: https://github.com/cai-zl/sql-runner.git

#### 使用说明

1. go run main.go run --help
2. go run main.go run -H 127.0.0.1 -P 5432 -u postgres -p 123456 -d carInfo -s 'show tables' -S 2 -D pgsql
