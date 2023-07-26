package service

import (
	"rgo/dbx/mysqlx"
	"rgo/logx"
	"rgo/web"
)

type Config struct {
	Name string
	Log  logx.LogConf

	Web   web.Config
	Mysql mysqlx.Config
}
