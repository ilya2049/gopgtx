package pg

import (
	"database/sql"
	"fmt"
	"log"

	// Postgres driver.
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func DefaultConfig() Config {
	return Config{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "test-user",
		Password: "password",
		Database: "test-db",
	}
}

func NewConnection(cfg Config) (*sql.DB, func(), error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open a postgres connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		log.Println()

		return nil, nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	log.Println("postgres: connection established")

	return db, func() {
		if err := db.Close(); err != nil {
			log.Println("postgres: failed to close postgres connection pull: " + err.Error())

			return
		}

		log.Println("postgres: connection closed")
	}, nil
}
