package api

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

const (
	MysqlDb = "mysql"
	PgsqlDb = "pgsql"
)

var (
	SqlInfoCache = map[string]func(info *SqlInfo) Runner{}
)
