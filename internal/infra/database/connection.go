package database

import (
	"database/sql"

	cnfg "github.com/LGuilhermeMoreira/url-shortener/config"
	_ "github.com/lib/pq"
)

func NewConnection(c *cnfg.Config) (*sql.DB, error) {
	return sql.Open(c.DBDriver, c.DBUri)
}
