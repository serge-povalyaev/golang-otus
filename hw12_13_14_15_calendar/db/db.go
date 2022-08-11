package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func CreateDSN(host, port, user, pass, name string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, name)
}

func Connect(dsn string) *sqlx.DB {
	connection, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logrus.Fatalln(err)
	}

	return connection
}
