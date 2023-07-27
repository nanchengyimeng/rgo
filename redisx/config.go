package redisx

// SingleConfig
// @Description: 单节点redis配置
type SingleConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// ClusterConfig
// @Description: 集群配置
type ClusterConfig struct {
	Addr     []string `yaml:"addr"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}
