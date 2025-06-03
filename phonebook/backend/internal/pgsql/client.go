package pgsql

import (
	"errors"

	"github.com/gocraft/dbr/v2"
)

var ErrNoRows = errors.New("err_no_rows")

type Client struct {
	pgdb *dbr.Connection
}

func NewClient(conn *dbr.Connection) *Client {
	return &Client{
		pgdb: conn,
	}
}

func (c *Client) GetSession() *dbr.Session {
	return c.pgdb.NewSession(nil)
}
