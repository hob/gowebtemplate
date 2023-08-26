package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"spillane.farm/gowebtemplate/internal/system"
)

var mdb *sql.DB

func Init(config system.Config) {
	logrus.Info("Initializing MySQL")
	d, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", config.MySqlUsername, config.MySqlPassword, config.MySqlHost, config.MySqlPath))
	if err != nil {
		logrus.WithError(err).Error("failed to initialize MySQL")
		panic(err.Error())
	}
	mdb = d
}

func Close() {
	mdb.Close()
}

func Query(query string) (*sql.Rows, error) {
	return mdb.Query(query)
}

func QueryRow(query string) *sql.Row {
	return mdb.QueryRow(query)
}
