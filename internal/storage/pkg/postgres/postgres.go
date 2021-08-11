package postgres

import (
	"fmt"

	"github.com/dmartzol/api-template/pkg/environment"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	dbport   = "DBPORT"
	dbuser   = "PGUSER"
	dbpass   = "PGPASSWORD"
	hostname = "PGHOST"
	dbname   = "PGDATABASE"
)

// DB represents the database
type DB struct {
	Client *sqlx.DB
}

type DatabaseConfig struct {
	Name, User, Password, Host string
	Port                       int
}

func NewDBClient(dbname, username, hostname string) (*DB, error) {
	dbConfig := DatabaseConfig{
		Port:     environment.GetEnvInt(dbport, 5432),
		Password: environment.GetEnvString(dbpass, ""),
		Host:     hostname,
		User:     username,
		Name:     dbname,
	}

	dataSourceName := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	dataSourceName = fmt.Sprintf(dataSourceName, dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)
	database, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to database")
	}
	err = database.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "error pinging database")
	}
	return &DB{database}, nil
}
