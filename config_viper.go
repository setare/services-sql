package srvsql

import (
	"time"

	"github.com/ory/viper"
)

type viperConfig struct {
	prefix string
}

func NewViperConfig(prefix string) DBServiceConfig {
	return &viperConfig{
		prefix: prefix,
	}
}

func (cfg *viperConfig) key(k string) string {
	if cfg.prefix == "" {
		return k
	}
	return cfg.prefix + "." + k
}

func (cfg *viperConfig) GetDSN() string {
	return viper.GetString(cfg.key("db.dsn"))
}

func (cfg *viperConfig) GetMaxIdleConnections() int {
	return viper.GetInt(cfg.key("db.max_idle_connections"))
}

func (cfg *viperConfig) GetMaxOpenConns() int {
	return viper.GetInt(cfg.key("db.max_open_conns"))
}

func (cfg *viperConfig) GetConnMaxLifetime() time.Duration {
	return viper.GetDuration(cfg.key("db.conn_max_lifetime"))
}
