package db

import (
	"authapp/packages/config"
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/apex/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)


func getDBURL() string {
	userName := config.ConfigSettings[config.DB_USERNAME]
	password := config.ConfigSettings[config.DB_PASSWORD]
	serviceHost := config.ConfigSettings[config.DB_SERVER_HOST]
	database	:=config.ConfigSettings[config.DB_DATABASE]

	connString := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", userName, password, serviceHost, database)
	return connString
}
 
func ConnectDB() (*sql.DB, error) {
	db, _ := sql.Open("postgres", getDBURL())
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func MigrateDBTables(db *sql.DB, dbName string) error {
	log.Info("running db migrations, to disable set RUN_MIGRATION=false")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	dir, _ := filepath.Abs("../packages/db/migrations")

	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", dir), dbName, driver)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}