package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

const (
	maxIdleConnection = 10
	maxOpenConnection = 10
)

func CreateDBConnection(
	username string,
	password string,
	host string,
	port string,
	dbName string,
) (db *sqlx.DB, err error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		port,
		dbName,
	)
	db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Failed to connect database")
	} else {
		log.Info().
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Connected to database")
	}

	db.SetMaxIdleConns(maxIdleConnection)
	db.SetMaxOpenConns(maxOpenConnection)

	return
}
