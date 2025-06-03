package schema

import (
	"errors"
	"fmt"

	"github.com/backend-d/phonebook/pkg/utils"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func MigratePostgres(cfg PostgresConfig) error {
	err := migrateUp(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode), "/postgres")
	if err != nil {
		return err
	}

	return nil
}

func migrateUp(conn, dir string) error {
	path, err := utils.CustomPath("/schema")
	if err != nil {
		return err
	}

	m, err := migrate.New(
		"file:"+path+dir,
		conn,
	)
	if err != nil {
		return fmt.Errorf("create migration: %s", err)
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("load migration: %s", err)
		}
	}

	return nil
}
