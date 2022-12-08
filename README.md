# sql-runner

#### 介绍
一个简单的sql定时执行器,支持cron、time,可用数据库: mysql、pgsql

#### 获取方式

1. github地址: https://github.com/cai-zl/sql-runner.git

#### 使用说明

1. 帮助: ./sql-runner run --help
2. time运行: ./sql-runner run -H 127.0.0.1 -P 5432 -u postgres -p 123456 -d carInfo -s 'show tables' -S 2 -D pgsql
3. cron运行: ./sql-runner run -H 127.0.0.1 -P 5432 -u postgres -p 123456 -d carInfo -s 'show tables' -C '0 * * * *' -D pgsql
4. time后台运行: ./sql-runner run -H 127.0.0.1 -P 5432 -u postgres -p 123456 -d carInfo -s 'show tables' -S 2 -D pgsql &
5. cron后台运行: ./sql-runner run -H 127.0.0.1 -P 5432 -u postgres -p 123456 -d carInfo -s 'show tables' -C '0 * * * *' -D pgsql &
