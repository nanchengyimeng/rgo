package dbx

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"rgo/dbx/mysqlx"
	"time"
)

type Driver struct {
	config mysqlx.Config
	db     *gorm.DB
	sqlDb  *sql.DB
}

func NewDriver() *Driver {
	return &Driver{}
}

// ConnectMysql
// @Description: 连接mysql
// @receiver d
// @param conf 数据库连接配置
// @param mysqlConfig  mysql的连接配置
// @return *Db
func (d *Driver) ConnectMysql(conf mysqlx.Config, gormConfigs ...*gorm.Config) {
	d.config = conf

	var gormConfig gorm.Config
	if len(gormConfigs) > 0 {
		gormConfig = *gormConfigs[0]
	}

	//加载日志配置
	d.logger(&gormConfig)

	db, err := gorm.Open(mysql.Open(d.config.Dsn), &gormConfig)
	if err != nil {
		panic(err)
	}

	d.db = db
	d.sqlDb, _ = db.DB()
	d.loadConfig()
}

func (d *Driver) GetDB() *gorm.DB {
	return d.db
}

// loadConfig
// @Description: 用户配置加载
// @receiver d
func (d *Driver) loadConfig() {
	if d.config.MaxIdleConn > 0 {
		d.sqlDb.SetMaxIdleConns(d.config.MaxIdleConn)
	}

	if d.config.MaxOpenConn > 0 {
		d.sqlDb.SetMaxOpenConns(d.config.MaxOpenConn)
	}

	if d.config.ConnMaxLifetime > 0 {
		d.sqlDb.SetConnMaxLifetime(time.Duration(d.config.ConnMaxLifetime) * time.Second)
	}
}

func (d *Driver) logger(config *gorm.Config) {
	var writer io.Writer

	if len(d.config.Path) != 0 {
		if _, err := os.Stat(d.config.Path); err != nil {
			err := os.MkdirAll(d.config.Path, 0711)
			if err != nil {
				panic(err)
				return
			}
		}

		file, err := os.OpenFile(d.config.Path+"/"+"sql.log", os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err)
		}

		writer = file
	} else {
		writer = os.Stdout
	}

	config.Logger = logger.New(
		log.New(writer, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,    // Slow SQL threshold
			LogLevel:                  logger.Info,    // Log level
			IgnoreRecordNotFoundError: d.config.Debug, // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,          // Don't include params in the SQL log
			Colorful:                  false,          // Disable color
		},
	)
}
