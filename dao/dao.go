package dao

import (
	"crypto/tls"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hawkingrei/bms/config"
)

type DAO struct {
	mysql *sql.DB
}

func NewDAO(cfg *config.Configure) (*DAO, error) {
	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: cfg.ServerName,
	})
	db, err := sql.Open("mysql", cfg.DbURL)
	if err != nil {
		return nil, err
	}
	return &DAO{
		mysql: db,
	}, nil
}
