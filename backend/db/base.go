package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	C "github.com/word-extractor/word-extractor-apis/config"
	_ "github.com/lib/pq"
)

var PostgresConn *sql.DB

// Returns postgres connection URL
func GetPostgresURL() string {
	dbHost := C.Conf.PostgresHost
	dbPort := C.Conf.PostgresPort
	dbUser := C.Conf.PostgresUser
	dbPass := C.Conf.PostgresPassword
	dbName := C.Conf.PostgresDB

	if C.Conf.PostgresSSLMode == "disable" {
		return fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPass, dbName)
	} else {
		return fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s sslrootcert=%s",
			dbHost, dbPort, dbUser, dbPass, dbName, C.Conf.PostgresSSLMode, C.Conf.PostgresRootCertLoc)
	}
}

// Configure postgres pooling logic
func ConfigurePGConn() {
	pgMaxOpenConns := C.Conf.PostgresMaxOpenConns

	PostgresConn.SetMaxOpenConns(pgMaxOpenConns)

	pgMaxIdleConns := C.Conf.PostgresMaxIdleConns

	PostgresConn.SetMaxIdleConns(pgMaxIdleConns)

	pgMaxIdleTime := C.Conf.PostgresMaxIdleTime

	PostgresConn.SetConnMaxIdleTime(pgMaxIdleTime)
}

// Returns a DB configuration object from the ENV.
func GetPGOptions() *pg.Options {
	dbHost := C.Conf.PostgresHost
	dbPort := C.Conf.PostgresPort
	dbUser := C.Conf.PostgresUser
	dbPass := C.Conf.PostgresPassword
	dbName := C.Conf.PostgresDB

	return &pg.Options{
		Addr:        fmt.Sprintf("%v:%v", dbHost, dbPort),
		User:        dbUser,
		Password:    dbPass,
		Database:    dbName,
		PoolSize:    C.Conf.PostgresMaxOpenConns,
		IdleTimeout: C.Conf.PostgresMaxIdleTime,
	}
}

// Initializes/configures pg
func Init() error {
	db, err := sql.Open("postgres", GetPostgresURL())
	if err != nil {
		return err
	}

	PostgresConn = db

	ConfigurePGConn()

	return nil
}

// Creates and returns a new transaction
func PGTransaction(ctx context.Context) (*sql.Tx, error) {
	tx, err := PostgresConn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Clean up all open connection pools
func Close() {
	err := PostgresConn.Close()

	if err != nil {
		log.Fatalln("Error while trying to close the postgres DB connection!", err)
	}
}
