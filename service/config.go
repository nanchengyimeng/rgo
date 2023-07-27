package service

import (
	"github.com/nanchengyimeng/rgo/dbx/mysqlx"
	"github.com/nanchengyimeng/rgo/logx"
	"github.com/nanchengyimeng/rgo/web"
)

type Config struct {
	Name string
	Log  logx.LogConf

	Web   web.Config
	Mysql mysqlx.Config
}
