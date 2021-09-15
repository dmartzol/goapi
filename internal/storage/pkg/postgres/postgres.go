package postgres

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// DB represents the database
type DB struct {
	Client *sqlx.DB
}

type Config struct {
	Name, User, Password, Host string
	Port                       int
}

func new(config Config) (*DB, error) {
	dataSourceName := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	dataSourceName = fmt.Sprintf(dataSourceName, config.Host, config.Port, config.User, config.Password, config.Name)
	database, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to database")
	}
	return &DB{database}, nil
}

func NewWithWaitLoop(config Config) (*DB, error) {
	for {
		db, err := new(config)
		if err != nil {
			if strings.Contains(err.Error(), "connection refused") {
				time.Sleep(time.Second * 1)
				continue
			}
			return nil, errors.Wrap(err, "error connecting to database")
		}
		return db, nil
	}
}
