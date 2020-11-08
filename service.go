package srvsql

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/spf13/viper"
	"github.com/pkg/errors"
	"github.com/setare/services"
)

type DBServiceConfig interface {
	GetDSN() string
	GetMaxIdleConnections() int
	GetMaxOpenConns() int
	GetConnMaxLifetime() time.Duration
}

type DBService interface {
	services.Service
	services.Startable
	WithName(string) DBService
	WithConfig(DBServiceConfig) DBService
	DB() *sql.DB
}

type dbService struct {
	db     *sql.DB
	name   string
	config DBServiceConfig
}

func NewDBService() DBService {
	return &dbService{}
}

func (service *dbService) WithName(name string) DBService {
	service.name = name
	return service
}

func (service *dbService) WithConfig(config DBServiceConfig) DBService {
	service.config = config
	return service
}

func (service *dbService) DB() *sql.DB {
	return service.db
}

// Name will return a human identifiable name for this service. Ex: Postgresql Connection.
func (service *dbService) Name() string {
	return service.name
}

// Start will start the service in a blocking way.
//
// If the service is successfully started, `nil` should be returned. Otherwise, an error must be returned.
func (service *dbService) Start() error {
	if service.config == nil {
		service.config = NewViperConfig("")
	}
	dsn := service.config.GetDSN()

	u, err := url.Parse(dsn)
	if err != nil {
		return errors.Wrap(err, "invalid DSN")
	}

	db, err := sql.Open(u.Scheme, dsn)
	if err != nil {
		return err
	}

	if maxIdleConnections := service.config.GetMaxIdleConnections(); maxIdleConnections > 0 {
		db.SetMaxIdleConns(maxIdleConnections)
	}

	if maxOpenConns := service.config.GetMaxOpenConns(); maxOpenConns > 0 {
		db.SetMaxOpenConns(maxOpenConns)
	}

	if connMaxLifetime := viper.GetDuration(""); connMaxLifetime > 0 {
		db.SetConnMaxLifetime(connMaxLifetime)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return err
	}

	service.db = db
	return nil
}

// Stop will stop this service.
//
// For most implementations it will be blocking and should return only when the service finishes stopping.
//
// If the service is successfully stopped, `nil` should be returned. Otherwise, an error must be returned.
func (service *dbService) Stop() error {
	return service.db.Close()
}
