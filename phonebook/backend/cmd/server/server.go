package main

import (
	"fmt"

	"github.com/backend-d/phonebook/internal/pgsql"
	"github.com/backend-d/phonebook/internal/server"
	"github.com/backend-d/phonebook/schema"
	"github.com/caarlos0/env/v6"
	"github.com/gocraft/dbr/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("loading .env file...")
	if err := godotenv.Load(); err != nil {
		log.Errorf("failed to load .env file: %v", err)
	}

	log.Info("loading config...")
	cfg := NewConfig()
	if err := env.Parse(&cfg); err != nil {
		log.Warnf("failed to parse .env to config: %v", err)
	}

	log.Info("creating pgdb connection...")
	conn, err := createPostgresConnection(cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to make pg connection: %v", err)
	}

	if err := schema.MigratePostgres(schema.PostgresConfig{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.Database,
		SSLMode:  cfg.Postgres.SSLMode,
	}); err != nil {
		log.Errorf("postgres migrate error: %v", err)
	}

	log.Info("loading pgsql client...")
	pg := pgsql.NewClient(conn)

	s := server.New(
		&server.Config{},
		&server.Deps{
			PG: pg,
		},
	)

	if err := s.App.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func createPostgresConnection(cfg PostgresConfig) (*dbr.Connection, error) {
	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}

	cs := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.SSLMode,
	)

	conn, err := dbr.Open("postgres", cs, nil)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
