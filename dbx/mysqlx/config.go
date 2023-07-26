package mysqlx

type Config struct {
	Dsn string `yaml:"dsn"`

	MaxIdleConn     int `yaml:"maxIdleConn"`
	MaxOpenConn     int `yaml:"maxOpenConn"`
	ConnMaxLifetime int `yaml:"connMaxLifetime"`

	Debug bool `yaml:"debug"`

	Path string `yaml:"path"`
}
