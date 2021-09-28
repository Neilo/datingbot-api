package postgree

import (
	"errors"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var instance *sqlx.DB

//Init db instance
func Init(dbURL string) error {
	db, err := sqlx.Open("pgx", dbURL)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	instance = db
	return nil
}

//GetDB return instance of DataBase
func GetDB() (*sqlx.DB, error) {
	if instance == nil {
		return nil, errors.New("Bad Postgre Store")
	}
	return instance, nil
}

//CloseDB close connection for PostgreSQL
func CloseDB() {
	if instance == nil {
		return
	}
	err := instance.Close()
	if err != nil {
		log.Info("Error when close PostgreSQL: %v", err)
	}
	instance = nil
}
