package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

func Migrate() {
	driver, err := mysql.WithInstance(mdb, &mysql.Config{})
	if err != nil {
		logrus.WithError(err).Error("failed to create mysql driver")
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
	if err != nil {
		logrus.WithError(err).Error("failed to run migrations")
		panic(err)
	}
	m.Steps(1)
}
