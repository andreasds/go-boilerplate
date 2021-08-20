package database

import (
	"fmt"

	"github.com/andreasds/go-boilerplate/configs"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

const (
	maxIdleConnection = 10
	maxOpenConnection = 10
)

type MySQLConnection struct {
	Read  *sqlx.DB
	Write *sqlx.DB
}

func NewMySQLConnection(config *configs.Config) *MySQLConnection {
	connection := new(MySQLConnection)
	connection.Read = createDBConnection(
		"read",
		config.DB.MySQL.Read.User,
		config.DB.MySQL.Read.Password,
		config.DB.MySQL.Read.Host,
		fmt.Sprint(config.DB.MySQL.Read.Port),
		config.DB.MySQL.Read.Name,
	)
	connection.Write = createDBConnection(
		"write",
		config.DB.MySQL.Write.User,
		config.DB.MySQL.Write.Password,
		config.DB.MySQL.Write.Host,
		fmt.Sprint(config.DB.MySQL.Write.Port),
		config.DB.MySQL.Write.Name,
	)

	return connection
}

func createDBConnection(
	method string,
	username string,
	password string,
	host string,
	port string,
	dbName string,
) (db *sqlx.DB) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		port,
		dbName,
	)
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("method", method).
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Failed to connect database")
	} else {
		log.Info().
			Str("method", method).
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Connected to database")
	}

	db.SetMaxIdleConns(maxIdleConnection)
	db.SetMaxOpenConns(maxOpenConnection)

	return
}
