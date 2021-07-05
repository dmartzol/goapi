package postgres

import (
	"fmt"
	"os"

	"github.com/dmartzol/api-template/pkg/environment"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	dbport = "DBPORT"
	dbuser = "PGUSER"
	dbpass = "PGPASSWORD"
	dbhost = "PGHOST"
	dbname = "PGDATABASE"
)

// DB represents the database
type DB struct {
	*sqlx.DB
}

type databaseConfig struct {
	Name, User, Password, Host string
	Port                       int
}

func NewDBClient() (*DB, error) {
	dbConfig := databaseConfig{}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		return nil, errors.New("PGDATABASE environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		return nil, errors.New("PGUSER environment variable required but not set")
	}
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		return nil, errors.New("PGHOST environment variable required but not set")
	}
	dbConfig.Port = environment.GetEnvInt(dbport, 5432)
	dbConfig.Password = environment.GetEnvString(dbpass, "")
	dbConfig.Host = host
	dbConfig.User = user
	dbConfig.Name = name

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
