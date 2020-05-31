package srvsql

import (
	"time"
)

type DBServiceConfigMemory interface {
	DBServiceConfig
	DSN(string) DBServiceConfigMemory
	MaxIdleConnections(int) DBServiceConfigMemory
	MaxOpenConns(int) DBServiceConfigMemory
	ConnMaxLifetime(time.Duration) DBServiceConfigMemory
}

type memoryConfig struct {
	dsn                string
	maxIdleConnections int
	maxOpenConns       int
	connMaxLifetime    time.Duration
}

func NewMemoryConfig() DBServiceConfigMemory {
	return &memoryConfig{}
}

func (cfg *memoryConfig) DSN(value string) DBServiceConfigMemory {
	cfg.dsn = value
	return cfg
}

func (cfg *memoryConfig) MaxIdleConnections(value int) DBServiceConfigMemory {
	cfg.maxIdleConnections = value
	return cfg
}

func (cfg *memoryConfig) MaxOpenConns(value int) DBServiceConfigMemory {
	cfg.maxOpenConns = value
	return cfg
}

func (cfg *memoryConfig) ConnMaxLifetime(value time.Duration) DBServiceConfigMemory {
	cfg.connMaxLifetime = value
	return cfg
}

func (cfg *memoryConfig) GetDSN() string {
	return cfg.dsn
}

func (cfg *memoryConfig) GetMaxIdleConnections() int {
	return cfg.maxIdleConnections
}

func (cfg *memoryConfig) GetMaxOpenConns() int {
	return cfg.maxOpenConns
}

func (cfg *memoryConfig) GetConnMaxLifetime() time.Duration {
	return cfg.connMaxLifetime
}
