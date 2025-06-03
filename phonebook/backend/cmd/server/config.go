package main

func NewConfig() Config {
	return Config{
		Server:   ServerConfig{},
		Postgres: PostgresConfig{},
	}
}

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	Port string `env:"SERVER_PORT" envDefault:"8080"`
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"postgres_db"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER" envDefault:"admin"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"admin123"`
	Database string `env:"POSTGRES_DATABASE" envDefault:"postgres"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
}
